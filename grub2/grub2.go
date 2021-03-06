/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package grub2

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"pkg.deepin.io/lib/procfs"

	"pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/log"
)

const grubScriptFile = "/boot/grub/grub.cfg"

var logger *log.Logger

func SetLogger(v *log.Logger) {
	logger = v
}

//go:generate dbusutil-gen -type Grub2,Theme grub2.go theme.go
type Grub2 struct {
	service       *dbusutil.Service
	modifyManager *modifyManager
	entries       []Entry
	theme         *Theme
	PropsMu       sync.RWMutex
	// props:
	DefaultEntry string
	EnableTheme  bool
	Resolution   string
	Timeout      uint32
	Updating     bool

	methods *struct {
		GetSimpleEntryTitles    func() `out:"titles"` // ([]string, *dbus.Error) {
		GetAvailableResolutions func() `out:"modeJSON"`
		SetDefaultEntry         func() `in:"entry"`
		SetEnableTheme          func() `in:"enabled"`
		SetResolution           func() `in:"resolution"`
		SetTimeout              func() `in:"timeout"`
	}
}

// return -1 for failed
func (g *Grub2) defaultEntryStr2Idx(str string) int {
	entriesLv1 := g.getEntryTitlesLv1()
	return getStringIndexInArray(str, entriesLv1)
}

func (g *Grub2) defaultEntryIdx2Str(idx int) (string, error) {
	entriesLv1 := g.getEntryTitlesLv1()
	length := len(entriesLv1)
	if length == 0 {
		return "", errors.New("no entry")
	}
	if 0 <= idx && idx < length {
		return entriesLv1[idx], nil
	} else {
		return "", errors.New("index out of range")
	}
}

func (g *Grub2) applyParams(params map[string]string) {
	//timeout
	timeout := getTimeout(params)
	if timeout < 0 {
		timeout = 999
	}
	g.Timeout = uint32(timeout)

	// enable theme
	var enableTheme bool
	theme := getTheme(params)
	if theme != "" {
		enableTheme = true
	}
	g.EnableTheme = enableTheme

	// resolution
	g.Resolution = getGfxMode(params)

	// default entry
	defaultEntry := getDefaultEntry(params)

	defaultEntryIdx, err := strconv.Atoi(defaultEntry)
	if err == nil {
		// is a num
		g.DefaultEntry, _ = g.defaultEntryIdx2Str(defaultEntryIdx)
	} else {
		// not a num
		if defaultEntry == "saved" {
			// TODO saved
			g.DefaultEntry, _ = g.defaultEntryIdx2Str(0)
		} else {
			g.DefaultEntry = defaultEntry
		}
	}
}

type modifyTask struct {
	paramsModifyFunc func(map[string]string)
	adjustTheme      bool
	adjustThemeLang  string
}

func getModifyTaskEnableTheme(enable bool, lang string) modifyTask {
	if enable {
		f := func(params map[string]string) {
			params[grubTheme] = quoteString(defaultGrubTheme)
			params[grubBackground] = quoteString(defaultGrubBackground)
		}
		return modifyTask{
			paramsModifyFunc: f,
			adjustTheme:      true,
			adjustThemeLang:  lang,
		}
	} else {
		f := func(params map[string]string) {
			delete(params, grubTheme)
			delete(params, grubBackground)
		}
		return modifyTask{
			paramsModifyFunc: f,
		}
	}
}

func getModifyTaskTimeout(timeout uint32) modifyTask {
	f := func(params map[string]string) {
		params[grubTimeout] = strconv.Itoa(int(timeout))
	}
	return modifyTask{
		paramsModifyFunc: f,
	}
}

func getModifyTaskResolution(val string, lang string) modifyTask {
	f := func(params map[string]string) {
		params[grubGfxMode] = quoteString(val)
	}
	return modifyTask{
		paramsModifyFunc: f,
		adjustTheme:      true,
		adjustThemeLang:  lang,
	}
}

func getModifyTaskDefaultEntry(idx int) modifyTask {
	f := func(params map[string]string) {
		params[grubDefault] = strconv.Itoa(idx)
	}
	return modifyTask{
		paramsModifyFunc: f,
	}
}

func New(service *dbusutil.Service) *Grub2 {
	g := &Grub2{
		service: service,
	}

	g.readEntries()

	params, err := loadGrubParams()
	if err != nil {
		logger.Warning(err)
	}

	g.applyParams(params)
	g.modifyManager = newModifyManager(func(running bool) {
		// state change callback
		g.setPropUpdating(running)
	})
	go g.modifyManager.loop()

	// init theme
	g.theme = NewTheme(g)

	jobLog, err := loadLog()
	if err != nil {
		if !os.IsNotExist(err) {
			logger.Warning(err)
		}
	}
	if jobLog != nil {
		if !jobLog.isJobDone(logJobMkConfig) {
			task := modifyTask{}

			if jobLog.hasJob(logJobAdjustTheme) &&
				!jobLog.isJobDone(logJobAdjustTheme) {
				task.adjustTheme = true
			}
			g.addModifyTask(task)
		}
	}

	return g
}

func (grub *Grub2) readEntries() (err error) {
	fileContent, err := ioutil.ReadFile(grubScriptFile)
	if err != nil {
		logger.Error(err)
		return
	}
	err = grub.parseEntries(string(fileContent))
	if err != nil {
		logger.Error(err)
		return
	}
	if len(grub.entries) == 0 {
		logger.Warning("there is no menu entry in %s", grubScriptFile)
	}
	return
}

func (grub *Grub2) resetEntries() {
	grub.entries = make([]Entry, 0)
}

// getAllEntriesLv1 return all entires titles in level one.
func (grub *Grub2) getEntryTitlesLv1() (entryTitles []string) {
	for _, entry := range grub.entries {
		if entry.parentSubMenu == nil {
			entryTitles = append(entryTitles, entry.getFullTitle())
		}
	}
	return
}

func (grub *Grub2) parseEntries(fileContent string) (err error) {
	grub.resetEntries()

	inMenuEntry := false
	level := 0
	numCount := make(map[int]int)
	numCount[0] = 0
	parentMenus := make([]*Entry, 0)
	parentMenus = append(parentMenus, nil)
	sl := bufio.NewScanner(strings.NewReader(fileContent))
	sl.Split(bufio.ScanLines)
	for sl.Scan() {
		line := sl.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "menuentry ") {
			if inMenuEntry {
				grub.resetEntries()
				err = fmt.Errorf("a 'menuentry' directive was detected inside the scope of a menuentry")
				return
			}
			title, ok := parseTitle(line)
			if ok {
				entry := Entry{MENUENTRY, title, numCount[level], parentMenus[len(parentMenus)-1]}
				grub.entries = append(grub.entries, entry)
				logger.Debugf("found entry: [%d] %s %s", level, strings.Repeat(" ", level*2), title)

				numCount[level]++
				inMenuEntry = true
				continue
			} else {
				grub.resetEntries()
				err = fmt.Errorf("parse entry title failed from: %q", line)
				return
			}
		} else if strings.HasPrefix(line, "submenu ") {
			if inMenuEntry {
				grub.resetEntries()
				err = fmt.Errorf("a 'submenu' directive was detected inside the scope of a menuentry")
				return
			}
			title, ok := parseTitle(line)
			if ok {
				entry := Entry{SUBMENU, title, numCount[level], parentMenus[len(parentMenus)-1]}
				grub.entries = append(grub.entries, entry)
				parentMenus = append(parentMenus, &entry)
				logger.Debugf("found entry: [%d] %s %s", level, strings.Repeat(" ", level*2), title)

				level++
				numCount[level] = 0
				continue
			} else {
				grub.resetEntries()
				err = fmt.Errorf("parse entry title failed from: %q", line)
				return
			}
		} else if line == "}" {
			if inMenuEntry {
				inMenuEntry = false
			} else if level > 0 {
				level--

				// delete last parent submenu
				i := len(parentMenus) - 1
				copy(parentMenus[i:], parentMenus[i+1:])
				parentMenus[len(parentMenus)-1] = nil
				parentMenus = parentMenus[:len(parentMenus)-1]
			}
		}
	}
	err = sl.Err()
	if err != nil {
		return
	}
	return
}

var (
	entryRegexpSingleQuote = regexp.MustCompile(`^ *(menuentry|submenu) +'(.*?)'.*$`)
	entryRegexpDoubleQuote = regexp.MustCompile(`^ *(menuentry|submenu) +"(.*?)".*$`)
)

func parseTitle(line string) (string, bool) {
	line = strings.TrimLeftFunc(line, unicode.IsSpace)
	if entryRegexpSingleQuote.MatchString(line) {
		return entryRegexpSingleQuote.FindStringSubmatch(line)[2], true
	} else if entryRegexpDoubleQuote.MatchString(line) {
		return entryRegexpDoubleQuote.FindStringSubmatch(line)[2], true
	} else {
		return "", false
	}
}

func (g *Grub2) getScreenWidthHeight() (w, h uint16, err error) {
	return parseResolution(g.Resolution)
}

func (g *Grub2) canSafelyExit() bool {
	logger.Debug("call canSafelyExit")
	if g.modifyManager.IsRunning() {
		return false
	}
	return true
}

func (g *Grub2) checkAuth(sender dbus.Sender) error {
	if noCheckAuth {
		logger.Warning("check auth disabled")
		return nil
	}

	pid, err := g.service.GetConnPID(string(sender))
	isAuthorized, err := checkAuthWithPid(pid)
	if err != nil {
		return err
	}
	if !isAuthorized {
		return errAuthFailed
	}
	return nil
}

func (g *Grub2) addModifyTask(task modifyTask) {
	g.modifyManager.ch <- task
}

func (g *Grub2) getSenderLang(sender dbus.Sender) (string, error) {
	pid, err := g.service.GetConnPID(string(sender))
	if err != nil {
		return "", err
	}

	p := procfs.Process(pid)
	environ, err := p.Environ()
	if err != nil {
		return "", err
	}

	return environ.Get("LANG"), nil
}

type Resolution struct {
	width  int
	height int
}

type Resolutions []Resolution

func (v Resolutions) Len() int {
	return len(v)
}

func (v Resolutions) Less(i, j int) bool {
	a := v[i]
	b := v[j]

	if a.width < b.width {
		return true
	} else if a.width == b.width {
		return a.height < b.height
	}
	return false
}

func (v Resolutions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Resolutions) add(r Resolution) Resolutions {
	var found bool
	for _, r0 := range v {
		if r0 == r {
			found = true
			break
		}
	}
	if !found {
		return append(v, r)
	}
	return v
}

var vbeResolutionsCache Resolutions
var vbeResolutionsOnce sync.Once

func getVbeResolutions() Resolutions {
	vbeResolutionsOnce.Do(
		func() {
			var err error
			vbeResolutionsCache, err = getVbeResolutions0()
			if err != nil {
				logger.Warning(err)
			}
		})
	if len(vbeResolutionsCache) == 0 {
		return Resolutions{
			{1920, 1080},
			{1366, 768},
			{1024, 768},
			{800, 600},
		}
	}
	return vbeResolutionsCache
}

func getVbeResolutions0() (Resolutions, error) {
	output, err := exec.Command("hwinfo", "--vbe").Output()
	if err != nil {
		return nil, err
	}
	var resolutions Resolutions

	add := func(w, h []byte) {
		width, _ := strconv.Atoi(string(w))
		height, _ := strconv.Atoi(string(h))

		if height == 0 || width == 0 {
			return
		}
		resolutions = resolutions.add(Resolution{
			width:  width,
			height: height,
		})
	}

	//  Mode 0x0311: 640x480 (+1280), 16 bits
	var regMode = regexp.MustCompile(`Mode .*: (\d+)x(\d+) .*\d+ bits`)

	// Resolution: 720x400@70Hz
	var regResolution = regexp.MustCompile(`Resolution: (\d+)x(\d+)@\d+Hz`)

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Bytes()
		match := regMode.FindSubmatch(line)
		if len(match) == 3 {
			width := match[1]
			height := match[2]
			add(width, height)
			continue
		}

		match = regResolution.FindSubmatch(line)
		if len(match) == 3 {
			width := match[1]
			height := match[2]
			add(width, height)
			continue
		}
	}

	resolutions = resolutions.add(Resolution{1024, 768})
	sort.Sort(sort.Reverse(resolutions))
	return resolutions, nil
}

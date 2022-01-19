package profiles

import (
	"strings"

	helpers     "github.com/mgord9518/aisap/helpers"
	permissions "github.com/mgord9518/aisap/permissions"
)

// List of all profiles supported by aisal out of the box.
// Most of these have only been tested on my (Manjaro and Arch) systems, so they may not work correctly on yours
// If that is the case, please report the issue and any error messages you encounter so that I can try to fix them
var Profiles = map[string]permissions.AppImagePerms{
	"aranym jit": {
		Level: 3,
		Files:   []string{ "xdg-download:ro", "~/Games:ro", "~/Roms:ro" },
		Devices: []string{ "dri", "input" },
		Sockets: []string{ "x11", "audio", "network" },
	},
	"aranym mmu": {
		Level: 3,
		Files:   []string{ "xdg-download:ro", "~/Games:ro", "~/Roms:ro" },
		Devices: []string{ "dri", "input" },
		Sockets: []string{ "x11", "audio", "network" },
	},
	// Untested with Android device, left level 1 assuming it needs access to all
	// of `/dev`
	"apk editor studio": {
		Level: 1,
		Files:   []string{ "xdg-templates:rw", "xdg-download:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	// Any apps that require superuser can't be sandboxed in this way
	"balenaetcher": {
		Level: 0,
	},
	// Badlion (and others) might be able to get switched to level 2, so specify devices anyway
	// Proprietary and unofficial AppImages should be high priority to be sandboxed to the fullest extent
	"badlion client": {
		Level: 2,
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "audio", "network" },
	},
	"blender": {
		Level: 2,
		Devices: []string{ "dri" },
		Files:   []string{ "xdg-templates:rw", "xdg-documents:rw" },
		Sockets: []string{ "x11" },
	},
	"deemix-gui": {
		Level: 2,
		Devices: []string{ "dri" },
		Files:   []string{ "xdg-music:rw" },
		Sockets: []string{ "x11", "audio", "network" },
	},
	"desmume": {
		Level: 2,
		Files:   []string{ "xdg-download:rw", "~/Games:rw", "~/Roms:rw" },
		Devices: []string{ "dri", "input" },
		Sockets: []string{ "x11", "pulseaudio" },
	},
	"dolphin emulator": {
		Level: 2,
		Files:   []string{ "xdg-download:ro", "~/Games:ro", "~/Roms:ro" },
		Devices: []string{ "dri", "input" },
		Sockets: []string{ "x11", "alsa" },
	},
	"dust3d": {
		Level: 2,
		Devices: []string{ "dri" },
		Files:   []string{ "xdg-templates:rw", "xdg-documents:rw" },
		Sockets: []string{ "x11" },
	},
	"endless sky": {
		Level: 2,
		Devices: []string{ "dri" },
		Files:   []string{ "xdg-templates:rw", "xdg-documents:rw" },
		Sockets: []string{ "x11", "alsa" },
	},
	"firefox": {
		Level: 2,
		Files:   []string{ "xdg-download:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "pulseaudio", "network" },
	},
	"firefox beta": {
		Level: 2,
		Files:   []string{ "xdg-download:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "pulseaudio", "network" },
	},
	"firefox nightly": {
		Level: 2,
		Files:   []string{ "xdg-download:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "pulseaudio", "network" },
	},
	"fontforge": {
		Level: 2,
		Files:   []string{ "xdg-documents:rw", "~/.fonts:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	// Also requires `/etc/passwd`
	"freecad conda": {
		Level: 1,
		Files:   []string{ "xdg-documents:rw", "xdg-templates:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	"geometrize": {
		Level: 2,
		Files:   []string{ "xdg-pictures:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	"gnu image manipulation program": {
		Level: 1,
		Files:   []string{ "xdg-pictures:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	"inkscape": {
		Level: 2,
		Files:   []string{ "xdg-documents:rw", "xdg-pictures:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	// Network for plugins and syncing
	"joplin": {
		Level: 2,
		Files:   []string{ "xdg-documents:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "network" },
	},
	"krita": {
		Level: 2,
		Files:   []string{ "xdg-pictures:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	"leocad": {
		Level: 2,
		Files:   []string{ "xdg-documents:rw", "xdg-templates:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	"librewolf": {
		Level: 2,
		Files:   []string{ "xdg-download:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "wayland", "pulseaudio", "network" },
	},
	"linedancer": {
		Level: 3,
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	"liteide": {
		Level: 2,
		Files:   []string{ "xdg-documents:rw", "~/go:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	"microsoft edge": {
		Level: 2,
		Files:   []string{ "xdg-download:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "pulseaudio", "network" },
	},
	// Minecraft requires access to keyring in order to launch correctly,
	// until a fix is found Minecraft will have to be run without a sandbox
	"minecraft": {
		Level: 0,
	},
	// Minetest demonstrates that not all GUI apps need level 2 or lower
	// fully self-contained apps that don't use system fonts, etc. can be
	// run in level 3
	"minetest": {
		Level: 3,
		Devices: []string{ "dri", "input" },
		Sockets: []string{ "x11", "audio", "network" },
	},
	"mypaint": {
		Level: 2,
		Files:   []string{ "xdg-pictures:rw", "xdg-templates:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	// Write access given to because save files are stored in the same
	// directory as the rom
	"mgba": {
		Level: 2,
		Files:   []string{ "xdg-download:rw", "~/Games:rw", "~/Roms:rw" },
		Devices: []string{ "dri", "input" },
		Sockets: []string{ "x11", "audio" },
	},
	// Network needed for cloud service, and can run in level 2 if given
	// `/etc/passwd`
	// TODO: Provide a fake `/etc/passwd` when running in level 2 or 3
	"onlyoffice desktop editors": {
		Level: 1,
		Files:   []string{ "xdg-documents:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "audio", "network" },
	},
	"photogimp": {
		Level: 1,
		Files:   []string{ "xdg-pictures:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	"pixsrt": {
		Level: 2,
		Files:   []string{ "xdg-pictures:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	"powder toy": {
		Level: 2,
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "network" },
	},
	"ppsspp": {
		Level: 2,
		Files:   []string{ "xdg-download:ro", "~/Games:ro", "~/Roms:ro" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "audio" },
	},
	// Python given no extra permissions, but can easily be customized for
	// scripts that require more
	"python3.10.1": {
		Level: 3,
	},
	// Only partially tested (I don't have an RS acct) but title screen
	// works as intended
	"runelite": {
		Level: 1,
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "audio", "network" },
	},
	// Audio for notification sounds
	"sengi": {
		Level: 1,
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "audio", "network" },
	},
	// Link to device not tested
	"signal": {
		Level: 2,
		Devices: []string{ "dri" },
		Sockets: []string{ "x11", "network" },
	},
	"smallbasic": {
		Level: 2,
		Files:   []string{ "xdg-documents:rw" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	// I think it's an interesting idea to have a confined terminal
	// enviornment although it could also greatly hinder its usefullness
	// so I'd like to hear feedback
	// TODO: add more files but keep it isolated from the host system
	"station": {
		Level: 1,
		Devices: []string{ "dri" },
		Files:   []string{ "~/.config/nvim:ro", "~/.profile:ro",
						   "~/.bashrc:ro",      "~/.zshrc:ro",
						   "~/.viminfo:ro"},
		Sockets: []string{ "x11" },
	},
	// Untested with real equipment but launches
	// TODO: Properly test Subsurface
	"subsurface": {
		Level: 1,
		Files:   []string{ "xdg-documents:ro" },
		Devices: []string{ "dri" },
		Sockets: []string{ "x11" },
	},
	"supertuxkart": {
		Level: 2,
		Devices: []string{ "dri", "input" },
		Sockets: []string{ "x11", "audio", "network" },
	},
	"supertux 2": {
		Level: 2,
		Devices: []string{ "dri", "input" },
		Sockets: []string{ "x11", "audio", "network" },
	},
	"texstudio": {
		Level: 2,
		Devices: []string{ "dri" },
		Files:   []string{ "xdg-documents:rw", "xdg-templates:rw" },
		Sockets: []string{ "x11" },
	},
	"visual studio code": {
		Level: 2,
		Devices: []string{ "dri" },
		Files:   []string{ "xdg-documents:rw" },
		Sockets: []string{ "x11", "network" },
	},
	// Partially tested
	"yuzu": {
		Level: 2,
		Files:   []string{ "xdg-download:ro", "~/Games:ro", "~/Roms:ro" },
		Devices: []string{ "dri", "input" },
		Sockets: []string{ "x11", "audio" },
	},
}

func FromName(name string) *permissions.AppImagePerms {
	name = strings.ToLower(name)

	if p, present := Profiles[name]; present {
		p.Files = helpers.CleanFiles(p.Files)
		return &p
	} else {
		p.Level = -1
		return &p
	}
}

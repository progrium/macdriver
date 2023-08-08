package modules

import "strings"

// Module Objective-c module
type Module struct {
	Name     string   // module name
	Title    string   // readable name
	Package  string   // go package for this module
	Header   string   // objc header file
	Prefixes []string // symbol prefixes
}

func (m *Module) String() string {
	return m.Name
}

func Get(moduleName string) *Module {
	for _, module := range All {
		if strings.EqualFold(moduleName, module.Name) ||
			moduleName == module.Title ||
			moduleName == module.Package {
			return &module
		}
	}
	return nil
}

func TrimPrefix(symbolName string) string {
	for _, m := range All {
		for _, prefix := range m.Prefixes {
			if strings.HasPrefix(symbolName, prefix) {
				name := strings.TrimPrefix(symbolName, prefix)
				if strings.HasPrefix(symbolName, "k") {
					return "K" + name
				}
				return name
			}
		}
	}
	return symbolName
}

var All = []Module{
	{"objectivec", "Objective-C Runtime", "objc", "objc/runtime.h", []string{}},
	{"Foundation", "Foundation", "foundation", "Foundation/Foundation.h", []string{"NS"}},
	{"AppKit", "AppKit", "appkit", "Appkit/Appkit.h", []string{"NS"}},
	{"UIKit", "UIKit", "uikit", "UIKit/UIKit.h", []string{"NS"}},
	{"UniformTypeIdentifiers", "Uniform Type Identifiers", "uniformtypeidentifiers", "UniformTypeIdentifiers/UniformTypeIdentifiers.h", []string{"UT"}},
	{"WebKit", "WebKit", "webkit", "WebKit/WebKit.h", []string{"WK"}},
	{"FileProvider", "File Provider", "fileprovider", "FileProvider/FileProvider.h", []string{"NS"}},
	{"Quartz", "Quartz", "quartz", "Quartz/Quartz.h", []string{"IK", "kQC", "kQuartz", "QC", "IK_"}},
	{"SecurityInterface", "Security Interface", "securityinterface", "SecurityInterface/SecurityInterface.h", []string{"SF"}},
	{"IOBluetooth", "IOBluetooth", "iobluetooth", "IOBluetooth/IOBluetooth.h", []string{"kIOBluetooth", "kBluetooth", "IOBluetooth", "Bluetooth", "kFTS", "kOBEX"}},
	{"CoreGraphics", "Core Graphics", "coregraphics", "CoreGraphics/CoreGraphics.h", []string{"CG", "kCG"}},
	{"CoreFoundation", "Core Foundation", "corefoundation", "CoreFoundation/CoreFoundation.h", []string{"CF", "kCF"}},
	{"QuartzCore", "QuartzCore", "quartzcore", "QuartzCore/QuartzCore.h", []string{"kCA", "CA"}},
	{"Vision", "Vision", "vision", "Vision/Vision.h", []string{"VN"}},
	{"CoreML", "Core ML", "coreml", "CoreML/CoreML.h", []string{"ML"}},
}

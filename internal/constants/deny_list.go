package constants

// Skip these — everything else is potentially code
var SkipDirs = map[string]bool{
	// VCS
	".git": true, ".svn": true, ".hg": true,

	// Dependencies (vendor/installed packages)
	"node_modules": true, "vendor": true, ".venv": true,
	"venv": true, "__pycache__": true, ".tox": true,
	"site-packages": true, "dist-packages": true,
	"bower_components": true, "jspm_packages": true,
	"packages": true,              // C# NuGet restore folder
	".gradle":  true, ".m2": true, // Java build caches

	// Build outputs
	"dist": true, "build": true, "out": true, "target": true,
	"bin": true, "obj": true, ".next": true, ".nuxt": true,
	".output": true, "coverage": true, ".nyc_output": true,
	".pytest_cache": true,

	// IDE/Editor
	".idea": true, ".vscode": true, ".vs": true,
	".eclipse": true, ".settings": true,

	// OS
	".DS_Store": true, "Thumbs.db": true,
	// Test directories
	"test":          true,
	"tests":         true,
	"__tests__":     true, // Jest (JS/TS)
	"spec":          true,
	"specs":         true,
	"e2e":           true,
	"integration":   true,
	"fixtures":      true, // test fixture data
	"mocks":         true, // mock files
	"__mocks__":     true, // Jest mocks
	"testdata":      true, // Go test data
	"test_data":     true,
	"snapshots":     true, // Jest snapshots
	"__snapshots__": true,
}

var SkipExtensions = map[string]bool{
	// Compiled binaries & archives
	".exe": true, ".dll": true, ".so": true, ".dylib": true,
	".a": true, ".o": true, ".obj": true, ".lib": true,
	".pyc": true, ".pyo": true, ".class": true, ".jar": true,
	".war": true, ".ear": true, ".wasm": true,
	".zip": true, ".tar": true, ".gz": true, ".rar": true,
	".7z": true, ".bz2": true,

	// Media
	".png": true, ".jpg": true, ".jpeg": true, ".gif": true,
	".svg": true, ".ico": true, ".webp": true, ".bmp": true,
	".mp4": true, ".mp3": true, ".wav": true, ".mov": true,
	".avi": true, ".webm": true, ".ttf": true, ".woff": true,
	".woff2": true, ".eot": true,

	// Docs/text artifacts (not logic)
	".pdf": true, ".doc": true, ".docx": true,
	".xls": true, ".xlsx": true, ".ppt": true, ".pptx": true,

	// Lock files (dependency metadata, not logic)
	// handle by filename, not extension

	// Data / generated
	".min.js": true, ".min.css": true, // handle via filename check
	".map": true, ".snap": true,
	".pb": true, ".proto.go": true, // generated protobuf
	".pb.swift": true, ".pb.cc": true,

	// Certs / secrets artifacts
	".pem": true, ".crt": true, ".key": true, ".p12": true,

	// Database
	".sqlite": true, ".db": true,
}

// Ski p specific filenames regardless of extension
var SkipFilenames = map[string]bool{
	"package-lock.json": true,
	"yarn.lock":         true,
	"pnpm-lock.yaml":    true,
	"Cargo.lock":        true,
	"Gemfile.lock":      true,
	"poetry.lock":       true,
	"composer.lock":     true,
	"go.sum":            true, // go.mod is useful though!
	"flake.lock":        true,
	".DS_Store":         true,
	"Thumbs.db":         true,
	".gitignore":        true,
	".gitattributes":    true,
	".editorconfig":     true,
	"LICENSE":           true,
	"CHANGELOG.md":      true,
}

// Skip test files by filename pattern (check before sending to AI)
var SkipFilenamePatterns = []string{
	// Universal patterns
	"*.test.*",    // foo.test.js, foo.test.ts
	"*.spec.*",    // foo.spec.js, foo.spec.rb
	"*_test.*",    // foo_test.go, foo_test.py
	"*_spec.*",    // foo_spec.rb
	"test_*",      // test_foo.py (Python convention)
	"*.stories.*", // foo.stories.tsx (Storybook)
	"*.e2e.*",     // end-to-end tests
	"*.bench.*",   // benchmark files
}

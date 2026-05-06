package constants

// SkipDirs — directory names to skip entirely during walk
// matched against d.Name() (just the folder name, not full path)
var SkipDirs = map[string]bool{

	// ========================
	// VCS
	// ========================
	".git": true, // Git
	".svn": true, // Subversion
	".hg":  true, // Mercurial
	".bzr": true, // Bazaar

	// ========================
	// DEPENDENCIES
	// ========================

	// JavaScript / TypeScript
	"node_modules":     true,
	"bower_components": true,
	"jspm_packages":    true,

	// Python
	".venv":         true,
	"venv":          true,
	"env":           true,
	"__pycache__":   true,
	".tox":          true,
	"site-packages": true,
	"dist-packages": true,
	".pyenv":        true,
	"eggs":          true,
	".eggs":         true,

	// Go
	"vendor": true,

	// Java / Kotlin / Scala
	".gradle": true, // Gradle cache
	".m2":     true, // Maven local repo

	// .NET / C# / F#
	"packages":    true, // NuGet restore
	".nuget":      true,
	"paket-files": true, // Paket

	// Swift / Objective-C
	".build":   true, // Swift SPM
	"Pods":     true, // CocoaPods
	"Carthage": true, // Carthage

	// Dart / Flutter
	".pub-cache": true,
	".dart_tool": true,

	// Elixir
	"deps": true,

	// Haskell
	".cabal-sandbox": true,
	".stack-work":    true,

	// Lua
	"lua_modules": true,

	// R
	"renv":    true,
	"packrat": true,

	// Ruby
	".bundle": true,

	// ========================
	// BUILD OUTPUTS
	// ========================

	// Generic
	"dist":    true,
	"build":   true,
	"out":     true,
	"output":  true,
	"release": true,
	"debug":   true,
	"target":  true, // Rust / Maven / SBT
	"bin":     true,
	"obj":     true,

	// JavaScript / TypeScript
	".next":         true, // Next.js
	".nuxt":         true, // Nuxt.js
	".svelte-kit":   true, // SvelteKit
	".output":       true, // Nuxt 3
	".vite":         true, // Vite cache
	".parcel-cache": true, // Parcel
	".turbo":        true, // Turborepo
	".webpack":      true, // Webpack cache

	// Java / Kotlin / Scala
	"classes": true,

	// Python
	".pytest_cache": true,
	".mypy_cache":   true,
	".ruff_cache":   true,
	"htmlcov":       true, // coverage html output

	// Coverage (generic)
	"coverage":    true,
	".nyc_output": true, // JS coverage
	"lcov-report": true,

	// Elixir
	"_build": true,

	// CMake / C / C++
	"cmake-build-debug":   true,
	"cmake-build-release": true,
	"CMakeFiles":          true,

	// iOS / macOS / Xcode
	"DerivedData": true,

	// Android
	"generated": true,

	// ========================
	// TEST DIRECTORIES
	// ========================

	// Generic
	"test":        true,
	"tests":       true,
	"testing":     true,
	"e2e":         true,
	"integration": true,
	"fixtures":    true, // test fixture data
	"mocks":       true, // mock files
	"stubs":       true,
	"fakes":       true,

	// JavaScript / TypeScript
	"__tests__":     true, // Jest
	"__mocks__":     true, // Jest mocks
	"__snapshots__": true, // Jest snapshots
	"cypress":       true, // Cypress e2e
	"playwright":    true, // Playwright e2e

	// Ruby
	"spec":  true,
	"specs": true,

	// Go / Python
	"testdata":  true,
	"test_data": true,

	// Snapshots
	"snapshots": true,

	// ========================
	// IDE / EDITOR
	// ========================
	".idea":     true, // JetBrains (IntelliJ, GoLand, PyCharm etc)
	".vscode":   true, // VS Code
	".vs":       true, // Visual Studio
	".eclipse":  true, // Eclipse
	".settings": true, // Eclipse settings
	".metals":   true, // Scala Metals
	".bloop":    true, // Scala Bloop
	".ionide":   true, // F# Ionide

	// ========================
	// OS GENERATED
	// ========================
	"__MACOSX":                  true, // macOS zip artifacts
	".Spotlight-V100":           true, // macOS
	".Trashes":                  true, // macOS
	"$RECYCLE.BIN":              true, // Windows
	"System Volume Information": true, // Windows
}

// SkipExtensions — file extensions to skip
// matched against filepath.Ext(filename) — only the LAST extension e.g. ".go", ".js"
// NOTE: multi-part extensions like ".min.js", ".pb.go" go in SkipFilenamePatterns
var SkipExtensions = map[string]bool{

	// ========================
	// COMPILED BINARIES
	// ========================
	".exe":   true, // Windows executable
	".dll":   true, // Windows dynamic lib
	".so":    true, // Linux shared lib
	".dylib": true, // macOS dynamic lib
	".a":     true, // static lib (C/C++/Rust)
	".o":     true, // object file
	".obj":   true, // Windows object file
	".lib":   true, // Windows static lib
	".out":   true, // Unix executable output
	".elf":   true, // Linux ELF binary
	".bin":   true, // raw binary

	// JVM
	".class": true, // Java compiled bytecode
	".jar":   true, // Java archive
	".war":   true, // Java web archive
	".ear":   true, // Java enterprise archive
	".dex":   true, // Android Dalvik bytecode

	// .NET
	".pdb":   true, // .NET debug symbols
	".nupkg": true, // NuGet package
	".msil":  true, // .NET IL

	// Python
	".pyc": true, // Python compiled bytecode
	".pyo": true, // Python optimized bytecode
	".pyd": true, // Python extension module (Windows)

	// WebAssembly
	".wasm": true,

	// ========================
	// ARCHIVES
	// ========================
	".zip":  true,
	".tar":  true,
	".gz":   true,
	".rar":  true,
	".7z":   true,
	".bz2":  true,
	".xz":   true,
	".zst":  true, // Zstandard
	".tgz":  true, // tar.gz
	".tbz2": true, // tar.bz2
	".dmg":  true, // macOS disk image
	".iso":  true, // disk image
	".deb":  true, // Debian package
	".rpm":  true, // RPM package
	".apk":  true, // Android package
	".ipa":  true, // iOS package

	// ========================
	// MEDIA — Images
	// ========================
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
	".bmp":  true,
	".ico":  true,
	".svg":  true,
	".webp": true,
	".tiff": true,
	".tif":  true,
	".raw":  true,
	".heic": true,
	".heif": true,
	".avif": true,
	".psd":  true, // Photoshop
	".ai":   true, // Adobe Illustrator
	".xcf":  true, // GIMP
	// NOTE: .svg intentionally excluded — may contain inline logic in web projects

	// ========================
	// MEDIA — Video / Audio
	// ========================
	".mp4":  true,
	".mkv":  true,
	".mov":  true,
	".avi":  true,
	".webm": true,
	".flv":  true,
	".wmv":  true,
	".mp3":  true,
	".wav":  true,
	".flac": true,
	".aac":  true,
	".ogg":  true,
	".m4a":  true,

	// ========================
	// FONTS
	// ========================
	".ttf":   true,
	".otf":   true,
	".woff":  true,
	".woff2": true,
	".eot":   true,

	// ========================
	// DOCUMENTS
	// ========================
	".pdf":  true,
	".doc":  true,
	".docx": true,
	".xls":  true,
	".xlsx": true,
	".ppt":  true,
	".pptx": true,
	".odt":  true, // OpenDocument text
	".ods":  true, // OpenDocument spreadsheet
	".odp":  true, // OpenDocument presentation
	".rtf":  true,
	".epub": true,

	// ========================
	// DATA FILES
	// ========================
	".map":     true, // source maps (JS)
	".snap":    true, // test snapshots
	".pb":      true, // protobuf binary
	".parquet": true, // columnar data
	".arrow":   true, // Apache Arrow
	".avro":    true, // Avro data
	".feather": true, // Feather data
	".hdf5":    true, // HDF5
	".h5":      true, // HDF5
	".npy":     true, // NumPy array
	".npz":     true, // NumPy compressed
	".pkl":     true, // Python pickle
	".pickle":  true, // Python pickle
	".mat":     true, // MATLAB data file
	".rds":     true, // R data
	".rda":     true, // R data
	".sav":     true, // SPSS data
	".dta":     true, // Stata data

	// ========================
	// DATABASE
	// ========================
	".sqlite":  true,
	".sqlite3": true,
	".db":      true,
	".mdb":     true, // MS Access
	".accdb":   true, // MS Access
	".frm":     true, // MySQL table definition
	".ibd":     true, // InnoDB data file

	// ========================
	// CERTS / KEYS / SECRETS
	// ========================
	".pem":      true,
	".crt":      true,
	".cer":      true,
	".key":      true,
	".p12":      true,
	".pfx":      true,
	".jks":      true, // Java KeyStore
	".keystore": true,

	// ========================
	// LOGS
	// ========================
	".log":  true,
	".logs": true,

	// ========================
	// MACHINE LEARNING MODELS
	// ========================
	".pt":          true, // PyTorch model
	".pth":         true, // PyTorch checkpoint
	".ckpt":        true, // TensorFlow checkpoint
	".safetensors": true, // Hugging Face model
	".onnx":        true, // ONNX model
	".tflite":      true, // TensorFlow Lite
	".mlmodel":     true, // Apple CoreML
	".joblib":      true, // scikit-learn model
}

// SkipFilenames — exact filenames to skip regardless of extension
// matched against filepath.Base(path) i.e. just the filename
var SkipFilenames = map[string]bool{

	// ========================
	// LOCK FILES — JavaScript / Node
	// ========================
	"package-lock.json": true, // npm
	"yarn.lock":         true, // Yarn
	"pnpm-lock.yaml":    true, // pnpm
	"bun.lockb":         true, // Bun
	"shrinkwrap.json":   true, // npm shrinkwrap

	// ========================
	// LOCK FILES — Other Languages
	// ========================
	"Cargo.lock":         true, // Rust
	"Gemfile.lock":       true, // Ruby
	"poetry.lock":        true, // Python Poetry
	"Pipfile.lock":       true, // Python Pipenv
	"pdm.lock":           true, // Python PDM
	"uv.lock":            true, // Python uv
	"composer.lock":      true, // PHP
	"go.sum":             true, // Go (go.mod is kept — has useful dep info)
	"flake.lock":         true, // Nix
	"mix.lock":           true, // Elixir
	"pubspec.lock":       true, // Dart / Flutter
	"Package.resolved":   true, // Swift SPM
	"packages.lock.json": true, // .NET NuGet
	"paket.lock":         true, // .NET Paket
	"gradle.lockfile":    true, // Java Gradle
	"ivy.xml":            true, // Java Ivy (auto-generated)
	"shard.lock":         true, // Crystal
	"deno.lock":          true, // Deno

	// ========================
	// GIT
	// ========================
	".gitignore":     true,
	".gitattributes": true,
	".gitmodules":    true,
	".gitkeep":       true,
	".mailmap":       true, // Git author mapping

	// ========================
	// OS GENERATED
	// ========================
	".DS_Store":   true, // macOS
	"Thumbs.db":   true, // Windows
	"desktop.ini": true, // Windows
	".directory":  true, // Linux KDE
	".localized":  true, // macOS

	// ========================
	// EDITOR / FORMATTER CONFIG
	// ========================
	".editorconfig":    true, // universal editor config
	".prettierrc":      true, // JS/TS Prettier formatter
	".prettierignore":  true,
	".eslintignore":    true, // JS/TS ESLint
	".stylelintignore": true,
	".npmignore":       true,
	".dockerignore":    true,
	".jshintrc":        true,
	".babelrc":         true,
	".browserslistrc":  true,
	".nvmrc":           true, // Node version (nvm)
	".node-version":    true, // Node version
	".ruby-version":    true, // Ruby version
	".python-version":  true, // Python version (pyenv)
	".tool-versions":   true, // asdf version manager

	// ========================
	// LINTER / FORMATTER CONFIGS
	// ========================
	".pylintrc":         true, // Python pylint
	".flake8":           true, // Python flake8
	".rubocop.yml":      true, // Ruby RuboCop
	".rubocop_todo.yml": true, // Ruby RuboCop auto-generated
	"phpcs.xml":         true, // PHP CodeSniffer
	"phpmd.xml":         true, // PHP Mess Detector
	"phpstan.neon":      true, // PHP PHPStan
	"psalm.xml":         true, // PHP Psalm
	"checkstyle.xml":    true, // Java Checkstyle
	"pmd.xml":           true, // Java PMD
	"detekt.yml":        true, // Kotlin Detekt
	"scalafmt.conf":     true, // Scala Scalafmt
	"rustfmt.toml":      true, // Rust rustfmt
	".golangci.yml":     true, // Go golangci-lint
	".golangci.yaml":    true, // Go golangci-lint
	".swiftlint.yml":    true, // Swift SwiftLint
	".clang-format":     true, // C/C++ clang-format
	".clang-tidy":       true, // C/C++ clang-tidy
	"stylecop.json":     true, // C# StyleCop
	".hadolint.yaml":    true, // Dockerfile linter

	// ========================
	// DOCS / LEGAL
	// ========================
	"LICENSE":            true,
	"LICENSE.md":         true,
	"LICENSE.txt":        true,
	"LICENCE":            true, // British spelling
	"LICENCE.md":         true,
	"CHANGELOG.md":       true,
	"CHANGELOG.txt":      true,
	"CHANGELOG.rst":      true,
	"CONTRIBUTING.md":    true,
	"CONTRIBUTING.rst":   true,
	"CODE_OF_CONDUCT.md": true,
	"SECURITY.md":        true,
	"SUPPORT.md":         true,
	"AUTHORS":            true,
	"AUTHORS.md":         true,
	"CONTRIBUTORS":       true,
	"CONTRIBUTORS.md":    true,
	"COPYING":            true, // GPL license file
	"NOTICE":             true, // Apache license notice
	"PATENTS":            true,
	"TRADEMARKS":         true,

	// ========================
	// PACKAGE MANAGER / CI MISC CONFIG
	// ========================
	".npmrc":               true, // npm config
	".yarnrc":              true, // Yarn config
	".yarnrc.yml":          true, // Yarn Berry config
	".pnpmfile.cjs":        true, // pnpm hooks
	"bunfig.toml":          true, // Bun config
	".huskyrc":             true, // Git hooks (Husky)
	".huskyrc.json":        true,
	".huskyrc.yaml":        true,
	".lintstagedrc":        true, // lint-staged
	"commitlint.config.js": true, // commitlint
	".commitlintrc":        true,
	".czrc":                true, // commitizen
	"renovate.json":        true, // Renovate dependency bot
	".renovaterc":          true,
	"dependabot.yml":       true, // GitHub Dependabot
}

// SkipFilenamePatterns — glob patterns matched against filename (filepath.Base)
// used for multi-part extensions and naming conventions not catchable by exact match
var SkipFilenamePatterns = []string{

	// ========================
	// TEST FILES
	// ========================
	"*.test.*",    // foo.test.js, foo.test.ts (JS/TS Jest)
	"*.spec.*",    // foo.spec.js, foo.spec.rb
	"*_test.*",    // foo_test.go, foo_test.py (Go / Python)
	"*_test",      // foo_test (no extension, rare Go binaries)
	"*_spec.*",    // foo_spec.rb (Ruby RSpec)
	"test_*",      // test_foo.py (Python unittest convention)
	"*.stories.*", // foo.stories.tsx (Storybook)
	"*.e2e.*",     // end-to-end test files
	"*.bench.*",   // benchmark files
	"*_bench.*",   // Go benchmark files

	// ========================
	// GENERATED FILES
	// ========================
	"*.min.js",       // minified JavaScript
	"*.min.css",      // minified CSS
	"*.pb.go",        // protobuf generated Go
	"*.pb.swift",     // protobuf generated Swift
	"*.pb.cc",        // protobuf generated C++
	"*.pb.h",         // protobuf generated C++ header
	"*.generated.*",  // generic generated files
	"*_generated.*",  // Go generated (mock_generated.go)
	"*.gen.*",        // short-form generated
	"*_gen.*",        // short-form generated
	"zz_generated.*", // controller-gen (Kubernetes operators)
	"mock_*.go",      // Go mocks (mockery tool)
	"*_mock.*",       // mock files
	"*.mock.*",       // mock files
}

task :default => [:test]

desc "Run just tests no measurements"
task :test do
  sh "ginkgo -r -skipMeasurements ."
end

desc "Run tests including measure tests"
task :test_and_measure do
  sh "ginkgo -r ."
end

desc "Builds a static binary"
task :build do
  ENV["GOOS"] ||= `go env GOOS`.chomp
  ENV["GOARCH"] ||= `go env GOARCH`.chomp
  version = ENV["VERSION"] || "development"
  sha = `git rev-parse --short HEAD`.chomp
  date = Time.now.strftime("%F %T %z")

  Rake::Task["test"].execute

  sh "go build -o %s -ldflags '-X github.com/choria-io/go-choria/version.Version=%s -X github.com/choria-io/go-choria/version.SHA=%s -X \"github.com/choria-io/go-choria/version.BuildDate=%s\"'" % [
    output_name(version), version, sha, date
  ]
end

desc "Builds a Linux binary"
task :build_linux do
  ENV["GOOS"] = "linux"
  ENV["GOARCH"] = "amd64"

  Rake::Task["build"].execute
end

def output_name(version)
  return ENV["OUTPUT"] if ENV["OUTPUT"]

  if ENV["GOOS"] && ENV["GOARCH"]
    return "choria-%s-%s-%s" % [version, ENV["GOOS"].capitalize, ENV["GOARCH"]]
  else
    return "choria-%s" % [version]
  end
end

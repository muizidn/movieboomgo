Pod::Spec.new do |spec|

    spec.name         = "MovieBOOMThrift"
    spec.version      = "0.1.0"
    spec.summary      = "MovieBOOM thrift generated swift files"
    spec.homepage     = "https://github.com/muizidn/movieboomgo"
    spec.license      = { :type => "MIT", :file => "LICENSE" }
    spec.author       = { "Muhammad Muizzsuddin" => "muiz.idn@gmail.com" }
  
    spec.ios.deployment_target = "9.0"
    spec.swift_version = "5.0"
  
    spec.source        = { :git => "https://github.com/muizidn/movieboomgo.git", :tag => "#{spec.version}" }
    spec.source_files  = "thrift/gen-swift/**/*.{swift}"
    # spec.resources     = "thrift/**/*.thrift"
    
end

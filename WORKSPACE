workspace(name = "com_github_glennadjrussell_mdax")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "86a5fcfb73f443d9cbf8d387f689a0f545200ca1ba294234e3ef866c50a4c2c2",
    strip_prefix = "bazel-gazelle-b1e37463c2ef0997c5cf0e90cdf03e696caaed27",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/archive/b1e37463c2ef0997c5cf0e90cdf03e696caaed27.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "co_honnef_go_tools",
    importpath = "honnef.co/go/tools",
    sum = "h1:UoveltGrhghAA7ePc+e+QYDHXrBps2PqFZiHkGR/xK8=",
    version = "v0.0.1-2020.1.4",
)

go_repository(
    name = "com_github_afex_hystrix_go",
    importpath = "github.com/afex/hystrix-go",
    sum = "h1:rFw4nCn9iMW+Vajsk51NtYIcwSTkXr+JGrMd36kTDJw=",
    version = "v0.0.0-20180502004556-fa1af6a1f4f5",
)

go_repository(
    name = "com_github_ajstarks_svgo",
    importpath = "github.com/ajstarks/svgo",
    sum = "h1:wVe6/Ea46ZMeNkQjjBW6xcqyQA/j5e0D6GytH95g0gQ=",
    version = "v0.0.0-20180226025133-644b8db467af",
)

go_repository(
    name = "com_github_alecthomas_template",
    importpath = "github.com/alecthomas/template",
    sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
    version = "v0.0.0-20190718012654-fb15b899a751",
)

go_repository(
    name = "com_github_alecthomas_units",
    importpath = "github.com/alecthomas/units",
    sum = "h1:UQZhZ2O0vMHr2cI+DC1Mbh0TJxzA3RcLoMsFw+aXw7E=",
    version = "v0.0.0-20190924025748-f65c72e2690d",
)

go_repository(
    name = "com_github_antihax_optional",
    importpath = "github.com/antihax/optional",
    sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_armon_circbuf",
    importpath = "github.com/armon/circbuf",
    sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
    version = "v0.0.0-20150827004946-bbbad097214e",
)

go_repository(
    name = "com_github_armon_go_metrics",
    importpath = "github.com/armon/go-metrics",
    sum = "h1:O2sNqxBdvq8Eq5xmzljcYzAORli6RWCvEym4cJf9m18=",
    version = "v0.3.9",
)

go_repository(
    name = "com_github_armon_go_radix",
    importpath = "github.com/armon/go-radix",
    sum = "h1:F4z6KzEeeQIMeLFa97iZU6vupzoecKdU5TX24SNppXI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    sum = "h1:QN1nsY27ssD/JmW4s83qmSb+uL6DG4GmCDzjmJB4xUI=",
    version = "v1.40.45",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2",
    importpath = "github.com/aws/aws-sdk-go-v2",
    sum = "h1:ZbovGV/qo40nrOJ4q8G33AGICzaPI45FHQWJ9650pF4=",
    version = "v1.9.1",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_cloudwatch",
    importpath = "github.com/aws/aws-sdk-go-v2/service/cloudwatch",
    sum = "h1:w/fPGB0t5rWwA43mux4e9ozFSH5zF1moQemlA131PWc=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_aws_smithy_go",
    importpath = "github.com/aws/smithy-go",
    sum = "h1:AEwwwXQZtUwP5Mz506FeXXrKBe0jA8gVM+1gEcSRooc=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_benbjohnson_clock",
    importpath = "github.com/benbjohnson/clock",
    sum = "h1:Q92kusRqC1XV2MjkWETPvjJVqKetz1OzxZB7mHJLju8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_beorn7_perks",
    importpath = "github.com/beorn7/perks",
    sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_bgentry_speakeasy",
    importpath = "github.com/bgentry/speakeasy",
    sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_burntsushi_xgb",
    importpath = "github.com/BurntSushi/xgb",
    sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
    version = "v0.0.0-20160522181843-27f122750802",
)

go_repository(
    name = "com_github_casbin_casbin_v2",
    importpath = "github.com/casbin/casbin/v2",
    sum = "h1:/poEwPSovi4bTOcP752/CsTQiRz2xycyVKFG7GUhbDw=",
    version = "v2.37.0",
)

go_repository(
    name = "com_github_cenkalti_backoff_v4",
    importpath = "github.com/cenkalti/backoff/v4",
    sum = "h1:G2HAfAmvm/GcKan2oOQpBXOd2tT2G57ZnZGWa1PxPBQ=",
    version = "v4.1.1",
)

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    importpath = "github.com/census-instrumentation/opencensus-proto",
    sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_cespare_xxhash",
    importpath = "github.com/cespare/xxhash",
    sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_cespare_xxhash_v2",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:YRXhKfTDauu4ajMg1TPgFO5jnlC2HCbmLXMcTG5cbYE=",
    version = "v2.1.2",
)

go_repository(
    name = "com_github_chzyer_logex",
    importpath = "github.com/chzyer/logex",
    sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_chzyer_readline",
    importpath = "github.com/chzyer/readline",
    sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
    version = "v0.0.0-20180603132655-2972be24d48e",
)

go_repository(
    name = "com_github_chzyer_test",
    importpath = "github.com/chzyer/test",
    sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
    version = "v0.0.0-20180213035817-a1ea475d72b1",
)

go_repository(
    name = "com_github_circonus_labs_circonus_gometrics",
    importpath = "github.com/circonus-labs/circonus-gometrics",
    sum = "h1:C29Ae4G5GtYyYMm1aztcyj/J5ckgJm2zwdDajFbx1NY=",
    version = "v2.3.1+incompatible",
)

go_repository(
    name = "com_github_circonus_labs_circonusllhist",
    importpath = "github.com/circonus-labs/circonusllhist",
    sum = "h1:TJH+oke8D16535+jHExHj4nQvzlZrj7ug5D7I/orNUA=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_clbanning_mxj",
    importpath = "github.com/clbanning/mxj",
    sum = "h1:HuhwZtbyvyOw+3Z1AowPkU87JkJUSv751ELWaiTpj8I=",
    version = "v1.8.4",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_cncf_udpa_go",
    importpath = "github.com/cncf/udpa/go",
    sum = "h1:cqQfy1jclcSy/FwLjemeg3SR1yaINm74aQyupQ0Bl8M=",
    version = "v0.0.0-20201120205902-5459f2c99403",
)

go_repository(
    name = "com_github_cncf_xds_go",
    importpath = "github.com/cncf/xds/go",
    sum = "h1:OZmjad4L3H8ncOIR8rnb5MREYqG8ixi5+WbeUsquF0c=",
    version = "v0.0.0-20210312221358-fbca930ec8ed",
)

go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_coreos_go_systemd_v22",
    importpath = "github.com/coreos/go-systemd/v22",
    sum = "h1:D9/bQk5vlXQFZ6Kwuu6zaiXJ9oTPe68++AzAJc1DzSI=",
    version = "v22.3.2",
)

go_repository(
    name = "com_github_creack_pty",
    importpath = "github.com/creack/pty",
    sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
    version = "v1.1.9",
)

go_repository(
    name = "com_github_datadog_datadog_go",
    importpath = "github.com/DataDog/datadog-go",
    sum = "h1:qSG2N4FghB1He/r2mFrWKCaL7dXCilEuNEeAn20fdD4=",
    version = "v3.2.0+incompatible",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_dustin_go_humanize",
    importpath = "github.com/dustin/go-humanize",
    sum = "h1:VSnTsYCnlFHaM2/igO1h6X3HA71jcobQuxemgkq4zYo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_eapache_go_resiliency",
    importpath = "github.com/eapache/go-resiliency",
    sum = "h1:1NtRmCAqadE2FN4ZcN6g90TP3uk8cg9rn9eNK2197aU=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_eapache_go_xerial_snappy",
    importpath = "github.com/eapache/go-xerial-snappy",
    sum = "h1:YEetp8/yCZMuEPMUDHG0CW/brkkEp8mzqk2+ODEitlw=",
    version = "v0.0.0-20180814174437-776d5712da21",
)

go_repository(
    name = "com_github_eapache_queue",
    importpath = "github.com/eapache/queue",
    sum = "h1:YOEu7KNc61ntiQlcEeUIoDTJ2o8mQznoNvUhiigpIqc=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_edsrzf_mmap_go",
    importpath = "github.com/edsrzf/mmap-go",
    sum = "h1:CEBF7HpRnUCSJgGUb5h1Gm7e3VkmVDrR8lvWVLtrOFw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_envoyproxy_go_control_plane",
    importpath = "github.com/envoyproxy/go-control-plane",
    sum = "h1:dulLQAYQFYtG5MTplgNGHWuV2D+OBD+Z8lmDBmbLg+s=",
    version = "v0.9.9-0.20210512163311-63b5d3c536b0",
)

go_repository(
    name = "com_github_envoyproxy_protoc_gen_validate",
    importpath = "github.com/envoyproxy/protoc-gen-validate",
    sum = "h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_fatih_color",
    importpath = "github.com/fatih/color",
    sum = "h1:mRhaKNwANqRgUBGKmnI5ZxEk7QXmjQeCcuYFMX2bfcc=",
    version = "v1.12.0",
)

go_repository(
    name = "com_github_fogleman_gg",
    importpath = "github.com/fogleman/gg",
    sum = "h1:WXb3TSNmHp2vHoCroCIB1foO/yQ36swABL8aOVeDpgg=",
    version = "v1.2.1-0.20190220221249-0403632d5b90",
)

go_repository(
    name = "com_github_franela_goblin",
    importpath = "github.com/franela/goblin",
    sum = "h1:cZqz+yOJ/R64LcKjNQOdARott/jP7BnUQ9Ah7KaZCvw=",
    version = "v0.0.0-20210519012713-85d372ac71e2",
)

go_repository(
    name = "com_github_franela_goreq",
    importpath = "github.com/franela/goreq",
    sum = "h1:a9ENSRDFBUPkJ5lCgVZh26+ZbGyoVJG7yb5SSzF5H54=",
    version = "v0.0.0-20171204163338-bcd34c9993f8",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:hsms1Qyu0jgnwNXIxa+/V/PDsU6CfLf6CNO8H7IWoS4=",
    version = "v1.4.9",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_gl_glfw",
    importpath = "github.com/go-gl/glfw",
    sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
    version = "v0.0.0-20190409004039-e6da0acd62b1",
)

go_repository(
    name = "com_github_go_gl_glfw_v3_3_glfw",
    importpath = "github.com/go-gl/glfw/v3.3/glfw",
    sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
    version = "v0.0.0-20200222043503-6f7a984d4dc4",
)

go_repository(
    name = "com_github_go_kit_kit",
    importpath = "github.com/go-kit/kit",
    sum = "h1:e4o3o3IsBfAKQh5Qbbiqyfu97Ku7jrO/JbohvztANh4=",
    version = "v0.12.0",
)

go_repository(
    name = "com_github_go_kit_log",
    importpath = "github.com/go-kit/log",
    sum = "h1:7i2K3eKTos3Vc0enKCfnVcgHh2olr/MyfboYq7cAcFw=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_go_logfmt_logfmt",
    importpath = "github.com/go-logfmt/logfmt",
    sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_go_stack_stack",
    importpath = "github.com/go-stack/stack",
    sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_go_task_slim_sprig",
    importpath = "github.com/go-task/slim-sprig",
    sum = "h1:p104kn46Q8WdvHunIJ9dAyjPVtrBPhSr3KT2yUst43I=",
    version = "v0.0.0-20210107165309-348f09dbbbc0",
)

go_repository(
    name = "com_github_go_zookeeper_zk",
    importpath = "github.com/go-zookeeper/zk",
    sum = "h1:4mx0EYENAdX/B/rbunjlt5+4RTA/a9SMHBRuSKdGxPM=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_godbus_dbus_v5",
    importpath = "github.com/godbus/dbus/v5",
    sum = "h1:9349emZab16e7zQvpmsbtjc18ykshndd8y2PG3sgJbA=",
    version = "v5.0.4",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_golang_freetype",
    importpath = "github.com/golang/freetype",
    sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
    version = "v0.0.0-20170609003504-e2365dfdc4a0",
)

go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name = "com_github_golang_groupcache",
    importpath = "github.com/golang/groupcache",
    sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
    version = "v0.0.0-20210331224755-41bb18bfe9da",
)

go_repository(
    name = "com_github_golang_jwt_jwt_v4",
    importpath = "github.com/golang-jwt/jwt/v4",
    sum = "h1:RAqyYixv1p7uEnocuy8P1nru5wprCh/MH2BIlW5z5/o=",
    version = "v4.0.0",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    sum = "h1:l75CXGRSwbaYNpl/Z2X1XIIAMSCquvXgpVZDhwEIJsc=",
    version = "v1.4.4",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:ROPKBNFfQgOUMifHyP+KYbvpjbdoFNs+aK7DXlji0Tw=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    sum = "h1:fHPg5GQYlCeLIPB9BZqMVR5nR9A+IM5zcgeTdjMYmLA=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_google_btree",
    importpath = "github.com/google/btree",
    sum = "h1:0udJVsspx3VBr5FwtLhQQtuAsVc79tTq0ocGIPAU6qo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:BKbKCqvP6I+rmFHt06ZmyQtvB8xAkWdhFyr0ZUNZcxQ=",
    version = "v0.5.6",
)

go_repository(
    name = "com_github_google_gofuzz",
    importpath = "github.com/google/gofuzz",
    sum = "h1:A8PeW59pxE9IoFRqBp37U+mSNaQoZ46F1f0f863XSXw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_google_martian_v3",
    importpath = "github.com/google/martian/v3",
    sum = "h1:pMen7vLs8nvgEYhywH3KDWJIJTeEr2ULsVWHWYHQyBs=",
    version = "v3.0.0",
)

go_repository(
    name = "com_github_google_pprof",
    importpath = "github.com/google/pprof",
    sum = "h1:Ak8CrdlwwXwAZxzS66vgPt4U8yUZX7JwLvVR58FN5jM=",
    version = "v0.0.0-20200708004538-1a94d8640e99",
)

go_repository(
    name = "com_github_google_renameio",
    importpath = "github.com/google/renameio",
    sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:EVhdT+1Kseyi1/pUmXKaFxYsDNy9RQYkMWRH68J/W7Y=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    importpath = "github.com/googleapis/gax-go/v2",
    sum = "h1:sjZBwGj9Jlw33ImPtvFviGYvseOtDM7hkSKB7+Tv3SM=",
    version = "v2.0.5",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
    version = "v0.0.0-20181017120253-0766667cb4d1",
)

go_repository(
    name = "com_github_gorilla_context",
    importpath = "github.com/gorilla/context",
    sum = "h1:AWwleXJkX/nhcU9bZSnZoi3h/qGYqQAGhq6zZe/aQW8=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_gorilla_mux",
    importpath = "github.com/gorilla/mux",
    sum = "h1:Pgr17XVTNXAk3q/r4CpKzC5xBM/qW1uVLV+IhRZpIIk=",
    version = "v1.6.2",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_prometheus",
    importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
    sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_hashicorp_consul_api",
    importpath = "github.com/hashicorp/consul/api",
    sum = "h1:MwZJp86nlnL+6+W1Zly4JUuVn9YHhMggBirMpHGD7kw=",
    version = "v1.10.1",
)

go_repository(
    name = "com_github_hashicorp_consul_sdk",
    importpath = "github.com/hashicorp/consul/sdk",
    sum = "h1:OJtKBtEjboEZvG6AOUdh4Z1Zbyu0WcxQ0qatRrZHTVU=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_hashicorp_errwrap",
    importpath = "github.com/hashicorp/errwrap",
    sum = "h1:hLrqtEDnRye3+sgx6z4qVLNuviH3MR5aQ0ykNJa/UYA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_cleanhttp",
    importpath = "github.com/hashicorp/go-cleanhttp",
    sum = "h1:035FKYIWjmULyFRBKPs8TBQoi0x6d9G4xc9neXJWAZQ=",
    version = "v0.5.2",
)

go_repository(
    name = "com_github_hashicorp_go_hclog",
    importpath = "github.com/hashicorp/go-hclog",
    sum = "h1:K4ev2ib4LdQETX5cSZBG0DVLk1jwGqSPXBjdah3veNs=",
    version = "v0.16.2",
)

go_repository(
    name = "com_github_hashicorp_go_immutable_radix",
    importpath = "github.com/hashicorp/go-immutable-radix",
    sum = "h1:DKHmCUm2hRBK510BaiZlwvpD40f8bJFeZnpfm2KLowc=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_hashicorp_go_msgpack",
    importpath = "github.com/hashicorp/go-msgpack",
    sum = "h1:zKjpN5BK/P5lMYrLmBHdBULWbJ0XpYR+7NGzqkZzoD4=",
    version = "v0.5.3",
)

go_repository(
    name = "com_github_hashicorp_go_multierror",
    importpath = "github.com/hashicorp/go-multierror",
    sum = "h1:B9UzwGQJehnUY1yNrnwREHc3fGbC2xefo8g4TbElacI=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_hashicorp_go_retryablehttp",
    importpath = "github.com/hashicorp/go-retryablehttp",
    sum = "h1:QlWt0KvWT0lq8MFppF9tsJGF+ynG7ztc2KIPhzRGk7s=",
    version = "v0.5.3",
)

go_repository(
    name = "com_github_hashicorp_go_rootcerts",
    importpath = "github.com/hashicorp/go-rootcerts",
    sum = "h1:jzhAVGtqPKbwpyCPELlgNWhE1znq+qwJtW5Oi2viEzc=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_hashicorp_go_sockaddr",
    importpath = "github.com/hashicorp/go-sockaddr",
    sum = "h1:GeH6tui99pF4NJgfnhp+L6+FfobzVW3Ah46sLo0ICXs=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_syslog",
    importpath = "github.com/hashicorp/go-syslog",
    sum = "h1:KaodqZuhUoZereWVIYmpUgZysurB1kBLX2j0MwMrUAE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_uuid",
    importpath = "github.com/hashicorp/go-uuid",
    sum = "h1:fv1ep09latC32wFoVwnqcnKJGnMSdBanPczbHAYm1BE=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    importpath = "github.com/hashicorp/golang-lru",
    sum = "h1:YDjusn29QI/Das2iO9M0BHnIbxPeyuCHsjMW+lJfyTc=",
    version = "v0.5.4",
)

go_repository(
    name = "com_github_hashicorp_logutils",
    importpath = "github.com/hashicorp/logutils",
    sum = "h1:dLEQVugN8vlakKOUE3ihGLTZJRB4j+M2cdTm/ORI65Y=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_mdns",
    importpath = "github.com/hashicorp/mdns",
    sum = "h1:XFSOubp8KWB+Jd2PDyaX5xUd5bhSP/+pTDZVDMzZJM8=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_hashicorp_memberlist",
    importpath = "github.com/hashicorp/memberlist",
    sum = "h1:5+RffWKwqJ71YPu9mWsF7ZOscZmwfasdA8kbdC7AO2g=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_hashicorp_serf",
    importpath = "github.com/hashicorp/serf",
    sum = "h1:EBWvyu9tcRszt3Bxp3KNssBMP1KuHWyO51lz9+786iM=",
    version = "v0.9.5",
)

go_repository(
    name = "com_github_hdrhistogram_hdrhistogram_go",
    importpath = "github.com/HdrHistogram/hdrhistogram-go",
    sum = "h1:5IcZpTvzydCQeHzK4Ef/D5rrSqwxob0t8PQPMybUNFM=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hudl_fargo",
    importpath = "github.com/hudl/fargo",
    sum = "h1:ZDDILMbB37UlAVLlWcJ2Iz1XuahZZTDZfdCKeclfq2s=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_ianlancetaylor_demangle",
    importpath = "github.com/ianlancetaylor/demangle",
    sum = "h1:UDMh68UUwekSh5iP2OMhRRZJiiBccgV7axzUG8vi56c=",
    version = "v0.0.0-20181102032728-5e5cf60278f6",
)

go_repository(
    name = "com_github_influxdata_influxdb1_client",
    importpath = "github.com/influxdata/influxdb1-client",
    sum = "h1:HqW4xhhynfjrtEiiSGcQUd6vrK23iMam1FO8rI7mwig=",
    version = "v0.0.0-20200827194710-b269163b24ab",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    importpath = "github.com/jmespath/go-jmespath",
    sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath_internal_testify",
    importpath = "github.com/jmespath/go-jmespath/internal/testify",
    sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_jpillora_backoff",
    importpath = "github.com/jpillora/backoff",
    sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
    version = "v1.1.12",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    importpath = "github.com/jstemmer/go-junit-report",
    sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
    version = "v4.20.0+incompatible",
)

go_repository(
    name = "com_github_julienschmidt_httprouter",
    importpath = "github.com/julienschmidt/httprouter",
    sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_jung_kurt_gofpdf",
    importpath = "github.com/jung-kurt/gofpdf",
    sum = "h1:PJr+ZMXIecYc1Ey2zucXdR73SMBtgjPgwa31099IMv0=",
    version = "v1.0.3-0.20190309125859-24315acbbda5",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_klauspost_compress",
    importpath = "github.com/klauspost/compress",
    sum = "h1:P76CopJELS0TiO2mebmnzgWaajssP/EszplttgQxcgc=",
    version = "v1.13.6",
)

go_repository(
    name = "com_github_knetic_govaluate",
    importpath = "github.com/Knetic/govaluate",
    sum = "h1:1G1pk05UrOh0NlF1oeaaix1x8XzrfjIDK47TY0Zehcw=",
    version = "v3.0.1-0.20171022003610-9aa49832a739+incompatible",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_kr_logfmt",
    importpath = "github.com/kr/logfmt",
    sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
    version = "v0.0.0-20140226030751-b84e30acd515",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:s5hAObm+yFO5uHYt5dYjxi2rXrsnmRpJx4OYvIWUaQs=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    importpath = "github.com/mattn/go-colorable",
    sum = "h1:c1ghPdyEDarC70ftn0y+A/Ee++9zz8ljHG1b13eJ0s8=",
    version = "v0.1.8",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:yVuAays6BHfxijgZPzw+3Zlu5yQgKGP2/hcQbHb7S9Y=",
    version = "v0.0.14",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    sum = "h1:4hp9jkHxhMHkqkrB3Ix0jegS5sx/RkqARlsWZ6pIwiU=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_miekg_dns",
    importpath = "github.com/miekg/dns",
    sum = "h1:JKfpVSCB84vrAmHzyrsxB5NAr5kLoMXZArPSw7Qlgyg=",
    version = "v1.1.43",
)

go_repository(
    name = "com_github_minio_highwayhash",
    importpath = "github.com/minio/highwayhash",
    sum = "h1:Aak5U0nElisjDCfPSG79Tgzkn2gl66NxOMspRrKnA/g=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_mitchellh_cli",
    importpath = "github.com/mitchellh/cli",
    sum = "h1:tEElEatulEHDeedTxwckzyYMA5c86fbmNIUL1hBIiTg=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    importpath = "github.com/mitchellh/go-homedir",
    sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_mitchellh_go_testing_interface",
    importpath = "github.com/mitchellh/go-testing-interface",
    sum = "h1:fzU/JVNcaqHQEcVFAKeR41fkiLdIPrefOvVG1VZ96U0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    sum = "h1:6h7AQ0yhTcIsmFmnAwQls75jp2Gzs4iB8W7pjMO+rqo=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    importpath = "github.com/modern-go/concurrent",
    sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    importpath = "github.com/mwitkow/go-conntrack",
    sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
    version = "v0.0.0-20190716064945-2f068394615f",
)

go_repository(
    name = "com_github_nats_io_jwt",
    importpath = "github.com/nats-io/jwt",
    sum = "h1:w3GMTO969dFg+UOKTmmyuu7IGdusK+7Ytlt//OYH/uU=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_nats_io_jwt_v2",
    importpath = "github.com/nats-io/jwt/v2",
    sum = "h1:i/O6cmIsjpcQyWDYNcq2JyZ3/VTF8SJ4JWluI5OhpvI=",
    version = "v2.0.3",
)

go_repository(
    name = "com_github_nats_io_nats_go",
    importpath = "github.com/nats-io/nats.go",
    sum = "h1:+0ndxwUPz3CmQ2vjbXdkC1fo3FdiOQDim4gl3Mge8Qo=",
    version = "v1.12.1",
)

go_repository(
    name = "com_github_nats_io_nats_server_v2",
    importpath = "github.com/nats-io/nats-server/v2",
    sum = "h1:wsnVaaXH9VRSg+A2MVg5Q727/CqxnmPLGFQ3YZYKTQg=",
    version = "v2.5.0",
)

go_repository(
    name = "com_github_nats_io_nkeys",
    importpath = "github.com/nats-io/nkeys",
    sum = "h1:cgM5tL53EvYRU+2YLXIK0G2mJtK12Ft9oeooSZMA2G8=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_nats_io_nuid",
    importpath = "github.com/nats-io/nuid",
    sum = "h1:5iA8DT8V7q8WK2EScv2padNa/rTESc1KdnPw4TC2paw=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_niemeyer_pretty",
    importpath = "github.com/niemeyer/pretty",
    sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
    version = "v0.0.0-20200227124842-a10e7caefd8e",
)

go_repository(
    name = "com_github_nxadm_tail",
    importpath = "github.com/nxadm/tail",
    sum = "h1:nPr65rt6Y5JFSKQO7qToXr7pePgD6Gwiw05lkbyAQTE=",
    version = "v1.4.8",
)

go_repository(
    name = "com_github_oneofone_xxhash",
    importpath = "github.com/OneOfOne/xxhash",
    sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    sum = "h1:HFB2fbVIlhIfCfOW81bZFbiC/RvnpXSdhbF2/DJr134=",
    version = "v1.16.2",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    sum = "h1:7lLHu94wT9Ij0o6EWWclhu0aOh32VxhkwEJvzuWPeak=",
    version = "v1.13.0",
)

go_repository(
    name = "com_github_op_go_logging",
    importpath = "github.com/op/go-logging",
    sum = "h1:lDH9UUVJtmYCjyT0CI4q8xvlXPxeZ0gYCVvWbmPlp88=",
    version = "v0.0.0-20160315200505-970db520ece7",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    sum = "h1:uEJPy/1a5RIPAJ0Ov+OIO8OxWu77jEv+1B0VhjKrZUs=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_openzipkin_zipkin_go",
    importpath = "github.com/openzipkin/zipkin-go",
    sum = "h1:UwtQQx2pyPIgWYHRg+epgdx1/HnBQTgN3/oIYEJTQzU=",
    version = "v0.2.5",
)

go_repository(
    name = "com_github_pascaldekloe_goe",
    importpath = "github.com/pascaldekloe/goe",
    sum = "h1:cBOtyMzM9HTpWjXfbbunk26uA6nG3a8n06Wieeh0MwY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_performancecopilot_speed_v4",
    importpath = "github.com/performancecopilot/speed/v4",
    sum = "h1:VxEDCmdkfbQYDlcr/GC9YoN9PQ6p8ulk9xVsepYy9ZY=",
    version = "v4.0.0",
)

go_repository(
    name = "com_github_pierrec_lz4",
    importpath = "github.com/pierrec/lz4",
    sum = "h1:VGcrWe3yk6o+t7BdVNy5UDPWa4OZuDWtE1W1ZbS7Kyw=",
    version = "v1.0.2-0.20190131084431-473cd7ce01a1",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_pkg_profile",
    importpath = "github.com/pkg/profile",
    sum = "h1:F++O52m40owAmADcojzM+9gyjmMOY/T4oYJkgFDH8RE=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_posener_complete",
    importpath = "github.com/posener/complete",
    sum = "h1:NP0eAhjcjImqslEwo/1hq7gpajME0fTLTezBKDqfXqo=",
    version = "v1.2.3",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    sum = "h1:HNkLOAEQMIDv/K+04rukrLx6ch7msSRwf3/SASFAGtQ=",
    version = "v1.11.0",
)

go_repository(
    name = "com_github_prometheus_client_model",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:uq5h0d+GuxiXLJLNABMgp2qUWDPiLvgCzz2dUR+/W/M=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_prometheus_common",
    importpath = "github.com/prometheus/common",
    sum = "h1:JEkYlQnpzrzQFxi6gnukFPdQ+ac82oRhzMcIduJu/Ug=",
    version = "v0.30.0",
)

go_repository(
    name = "com_github_prometheus_procfs",
    importpath = "github.com/prometheus/procfs",
    sum = "h1:4jVXhlkAyzOScmCkXBTOLRLTz8EeU+eyjrwB/EPq0VU=",
    version = "v0.7.3",
)

go_repository(
    name = "com_github_rcrowley_go_metrics",
    importpath = "github.com/rcrowley/go-metrics",
    sum = "h1:9ZKAASQSHhDYGoxY8uLVpewe1GDZ2vu2Tr/vTdVAkFQ=",
    version = "v0.0.0-20181016184325-3113b8401b8a",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    importpath = "github.com/rogpeppe/fastuuid",
    sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    importpath = "github.com/rogpeppe/go-internal",
    sum = "h1:RR9dF3JtopPvtkroDZuVD7qquD0bnHlKSqaQhgwt8yk=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_ryanuber_columnize",
    importpath = "github.com/ryanuber/columnize",
    sum = "h1:UFr9zpz4xgTnIE5yIMtWAMngCdZ9p/+q6lTbgelo80M=",
    version = "v0.0.0-20160712163229-9b3edd62028f",
)

go_repository(
    name = "com_github_sean_seed",
    importpath = "github.com/sean-/seed",
    sum = "h1:nn5Wsu0esKSJiIVhscUtVbo7ada43DJhG55ua/hjS5I=",
    version = "v0.0.0-20170313163322-e2103e2c3529",
)

go_repository(
    name = "com_github_shopify_sarama",
    importpath = "github.com/Shopify/sarama",
    sum = "h1:9oksLxC6uxVPHPVYUmq6xhr1BOF/hHobWH2UzO67z1s=",
    version = "v1.19.0",
)

go_repository(
    name = "com_github_shopify_toxiproxy",
    importpath = "github.com/Shopify/toxiproxy",
    sum = "h1:TKdv8HiTLgE5wdJuEML90aBgNWsokNbMijUGhmcoBJc=",
    version = "v2.1.4+incompatible",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:dJKuHgqk1NNQlqoA6BTlM1Wf9DOH3NBjQyu0h9+AZZE=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    importpath = "github.com/smartystreets/assertions",
    sum = "h1:zE9ykElWQ6/NYmHa3jpm/yHnI4xSofP+UP6SpjHcSeM=",
    version = "v0.0.0-20180927180507-b2de0cb4f26d",
)

go_repository(
    name = "com_github_smartystreets_goconvey",
    importpath = "github.com/smartystreets/goconvey",
    sum = "h1:fv0U8FUIMPNf1L9lnHLvLhgicrIVChEkdzIKYqbNC9s=",
    version = "v1.6.4",
)

go_repository(
    name = "com_github_sony_gobreaker",
    importpath = "github.com/sony/gobreaker",
    sum = "h1:oMnRNZXX5j85zso6xCPRNPtmAycat+WcoKbklScLDgQ=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_spaolacci_murmur3",
    importpath = "github.com/spaolacci/murmur3",
    sum = "h1:qLC7fQah7D6K1B0ujays3HV9gkFtllcxhzImRR7ArPQ=",
    version = "v0.0.0-20180118202830-f09979ecbc72",
)

go_repository(
    name = "com_github_streadway_amqp",
    importpath = "github.com/streadway/amqp",
    sum = "h1:kuuDrUJFZL1QYL9hUNuCxNObNzB0bV/ZG5jV3RWAQgo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_streadway_handy",
    importpath = "github.com/streadway/handy",
    sum = "h1:mOtuXaRAbVZsxAHVdPR3IjfmN8T1h2iczJLynhLybf8=",
    version = "v0.0.0-20200128134331-0f66f006fb2e",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:2vfRuCMp5sSVIDSqO8oNnWJq7mPa6KVP3iPIwFBuy8A=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:nwc3DEeHmmLAfoZucVR881uASk0Mfjw8xYJ99tb5CcY=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_tv42_httpunix",
    importpath = "github.com/tv42/httpunix",
    sum = "h1:G3dpKMzFDjgEh2q1Z7zUUtKa8ViPtH+ocF0bE0g00O8=",
    version = "v0.0.0-20150427012821-b75d8614f926",
)

go_repository(
    name = "com_github_vividcortex_gohistogram",
    importpath = "github.com/VividCortex/gohistogram",
    sum = "h1:6+hBz+qvs0JOrrNhhmR7lFxo5sINxBCGXrdtl/UvroE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_yuin_goldmark",
    importpath = "github.com/yuin/goldmark",
    sum = "h1:dPmz1Snjq0kmkz159iL7S6WzdahUTHnHB5M56WFVifs=",
    version = "v1.3.5",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    sum = "h1:Dg9iHVQfrhq82rUNu9ZxUDrJLaxFUe/HlCVaLyRruq8=",
    version = "v0.65.0",
)

go_repository(
    name = "com_google_cloud_go_bigquery",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:PQcPefKFdaIzjQFbiyOgAqyx8q5djaE7x9Sqe712DPA=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_datastore",
    importpath = "cloud.google.com/go/datastore",
    sum = "h1:/May9ojXjRkPBNVrq+oWLqmWCkr4OU5uRY29bu0mRyQ=",
    version = "v1.1.0",
)

go_repository(
    name = "com_google_cloud_go_pubsub",
    importpath = "cloud.google.com/go/pubsub",
    sum = "h1:ukjixP1wl0LpnZ6LWtZJ0mX5tBmjp1f8Sqer8Z2OMUU=",
    version = "v1.3.1",
)

go_repository(
    name = "com_google_cloud_go_storage",
    importpath = "cloud.google.com/go/storage",
    sum = "h1:STgFzyU5/8miMl0//zKh2aQeTyeaUH3WN9bSUiJ09bA=",
    version = "v1.10.0",
)

go_repository(
    name = "com_shuralyov_dmitri_gpu_mtl",
    importpath = "dmitri.shuralyov.com/gpu/mtl",
    sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
    version = "v0.0.0-20190408044501-666a987793e9",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
    sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
    version = "v2.2.6",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:BLraFXnmrev5lT+xlilqcH8XK9/i0At2xKjWk4p6zsU=",
    version = "v1.0.0-20200227125254-8fa46927fb4f",
)

go_repository(
    name = "in_gopkg_errgo_v2",
    importpath = "gopkg.in/errgo.v2",
    sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
    version = "v2.1.0",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
    version = "v1.4.7",
)

go_repository(
    name = "in_gopkg_gcfg_v1",
    importpath = "gopkg.in/gcfg.v1",
    sum = "h1:m8OOJ4ccYHnx2f4gQwpno8nAX5OGOh7RLaaz0pj3Ogs=",
    version = "v1.2.3",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    importpath = "gopkg.in/tomb.v1",
    sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
    version = "v1.0.0-20141024135613-dd632973f1e7",
)

go_repository(
    name = "in_gopkg_warnings_v0",
    importpath = "gopkg.in/warnings.v0",
    sum = "h1:wFXVbFY8DY5/xOe1ECiWdKCzZlxgshcYVNkBHstARME=",
    version = "v0.1.2",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
    version = "v2.4.0",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:h8qDotaEPuJATrMmW04NCwg7v22aHH28wwpauUhK9Oo=",
    version = "v3.0.0-20210107192922-496545a6307b",
)

go_repository(
    name = "io_etcd_go_etcd_api_v3",
    importpath = "go.etcd.io/etcd/api/v3",
    sum = "h1:GsV3S+OfZEOCNXdtNkBSR7kgLobAa/SO6tCxRa0GAYw=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_client_pkg_v3",
    importpath = "go.etcd.io/etcd/client/pkg/v3",
    sum = "h1:2aQv6F436YnN7I4VbI8PPYrBhu+SmrTaADcf8Mi/6PU=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_client_v2",
    importpath = "go.etcd.io/etcd/client/v2",
    sum = "h1:ftQ0nOOHMcbMS3KIaDQ0g5Qcd6bhaBrQT6b89DfwLTs=",
    version = "v2.305.0",
)

go_repository(
    name = "io_etcd_go_etcd_client_v3",
    importpath = "go.etcd.io/etcd/client/v3",
    sum = "h1:62Eh0XOro+rDwkrypAGDfgmNh5Joq+z+W9HZdlXMzek=",
    version = "v3.5.0",
)

go_repository(
    name = "io_k8s_sigs_yaml",
    importpath = "sigs.k8s.io/yaml",
    sum = "h1:kr/MCeFWJWTwyaHoR9c8EjH9OumOmoF9YGiZd7lFm/Q=",
    version = "v1.2.0",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    sum = "h1:gqCw0LfLxScz8irSi8exQc7fyQ0fKQU/qnC/X8+V/1M=",
    version = "v0.23.0",
)

go_repository(
    name = "io_opentelemetry_go_proto_otlp",
    importpath = "go.opentelemetry.io/proto/otlp",
    sum = "h1:rwOQPCuKAKmwGKq2aVNnYIibI6wnV7EvzgfTCzcdGg8=",
    version = "v0.7.0",
)

go_repository(
    name = "io_rsc_binaryregexp",
    importpath = "rsc.io/binaryregexp",
    sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
    version = "v0.2.0",
)

go_repository(
    name = "io_rsc_pdf",
    importpath = "rsc.io/pdf",
    sum = "h1:k1MczvYDUvJBe93bYd7wrZLLUEcLZAuF824/I4e5Xr4=",
    version = "v0.1.1",
)

go_repository(
    name = "io_rsc_quote_v3",
    importpath = "rsc.io/quote/v3",
    sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
    version = "v3.1.0",
)

go_repository(
    name = "io_rsc_sampler",
    importpath = "rsc.io/sampler",
    sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
    version = "v1.3.0",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    sum = "h1:yfrXXP61wVuLb0vBcG6qaOoIoqYEzOQS8jum51jkv2w=",
    version = "v0.30.0",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:lMO5rYAqUxkmaj76jAkRUvt5JZgFymx/+Q5Mzfivuhc=",
    version = "v1.6.6",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:ysnBoUyeL/H6RCvNRhWHjKoDEmguI+mPU+qHgK8qv/w=",
    version = "v0.0.0-20210917145530-b395a37504d4",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:AGJ0Ih4mHjSeibYkFGh1dD9KJ/eOtZ93I6hoHhukQ5Q=",
    version = "v1.40.0",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:SnqbnDw1V7RiZcXPx5MEeqPv2s79L9i7BJUlG/+RurQ=",
    version = "v1.27.1",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:3erb+vDS8lU1sxfDHF4/hhWyaXnhIaO+7RgL4fDZORA=",
    version = "v0.0.0-20210915214749-c084706c2272",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:QE6XYQK6naiK1EPAe1g/ILLxN5RBoH5xkJk3CqlMI/Y=",
    version = "v0.0.0-20200224162631-6cc2880d07d6",
)

go_repository(
    name = "org_golang_x_image",
    importpath = "golang.org/x/image",
    sum = "h1:+qEpEAPhDZ1o0x3tHzZTQDArnOixOzGD9HUJfcg0mb4=",
    version = "v0.0.0-20190802002840-cff245a6509b",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:VLliZ0d+/avPrXXH+OakdXhpJuEoBZuwh1m2j7U6Iug=",
    version = "v0.0.0-20210508222113-6edffad5e616",
)

go_repository(
    name = "org_golang_x_mobile",
    importpath = "golang.org/x/mobile",
    sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
    version = "v0.0.0-20190719004257-d2bd2a29d028",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:Gz96sIWK3OalVv/I/qNygP42zyoKp3xptRVCWRFEBvo=",
    version = "v0.4.2",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:R150MpwJIv1MpS0N/pc+NhTM8ajzvlmxlY5OYsrevXQ=",
    version = "v0.0.0-20210917221730-978cfadd31cf",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:pkQiBZBvdos9qq4wBAHqlzuZHEXo07pqV06ef90u1WI=",
    version = "v0.0.0-20210514164344-f6687ab2804c",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:5KslGYwFpkhGh+Q16bwMP3cOontH8FOep7tGV86Y7SQ=",
    version = "v0.0.0-20210220032951-036812b2e83c",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:J27LZFQBFoihqXoegpscI10HpjZ7B5WQLLKL2FZXQKw=",
    version = "v0.0.0-20210917161153-d61c044b1678",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:v+OssWQX+hTHEmOBgwxdZxK4zHq3yOs8F9J7mk0PY8E=",
    version = "v0.0.0-20201126162022-7de9c90e9dd1",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:olpwvP2KacW1ZWvsR7uQhoyTYvKAupfQrRGBFM352Gk=",
    version = "v0.3.7",
)

go_repository(
    name = "org_golang_x_time",
    importpath = "golang.org/x/time",
    sum = "h1:7zkz7BUtwNFFqcowJ+RIgu2MaV/MapERkDIy+mwPyjs=",
    version = "v0.0.0-20210723032227-1f47c861a9ac",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:ouewzE6p+/VEB31YYnTbEJdi8pFqKp4P4n85vwo3DHA=",
    version = "v0.1.5",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_repository(
    name = "org_gonum_v1_gonum",
    importpath = "gonum.org/v1/gonum",
    sum = "h1:CCXrcPKiGGotvnN6jfUsKk4rRqm7q09/YbKb5xCEvtM=",
    version = "v0.8.2",
)

go_repository(
    name = "org_gonum_v1_netlib",
    importpath = "gonum.org/v1/netlib",
    sum = "h1:OE9mWmgKkjJyEmDAAtGMPjXu+YNeGvK9VTSHY6+Qihc=",
    version = "v0.0.0-20190313105609-8cb42192e0e0",
)

go_repository(
    name = "org_gonum_v1_plot",
    importpath = "gonum.org/v1/plot",
    sum = "h1:Qh4dB5D/WpoUUp3lSod7qgoyEHbDGPUWjIbnqdqqe1k=",
    version = "v0.0.0-20190515093506-e2840ee46a6b",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    sum = "h1:ECmE8Bn/WFTYwEW/bpKD3M8VtR/zQVbavAoalC1PYyE=",
    version = "v1.9.0",
)

go_repository(
    name = "org_uber_go_goleak",
    importpath = "go.uber.org/goleak",
    sum = "h1:sHOAIxRGBp443oHZIPB+HsUGaksVCXVQENPxwTfQdH4=",
    version = "v1.1.11-0.20210813005559-691160354723",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    sum = "h1:zaiO/rmgFjbmCXdSYJWQcdvOCsthmdaHfr3Gm2Kx4Ec=",
    version = "v1.7.0",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    sum = "h1:ue41HOKd1vGURxrmeKIgELGb3jPW9DMUDGtsinblHwI=",
    version = "v1.19.1",
)

go_rules_dependencies()

go_register_toolchains(version = "1.16.5")

gazelle_dependencies()

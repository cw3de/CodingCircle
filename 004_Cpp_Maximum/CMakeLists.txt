cmake_minimum_required(VERSION 3.19)
project(CodingCircle004 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

include(FetchContent)
FetchContent_Declare(
  googletest
  GIT_REPOSITORY https://github.com/google/googletest.git
  GIT_TAG        03597a01ee50ed33e9dfd640b249b4be3799d395 # release-1.14.0
)

FetchContent_MakeAvailable(googletest)

enable_testing()

add_executable(
    RunningMaximum
    TestRunningMaximum.cpp
)

target_link_libraries(
    RunningMaximum
    GTest::gtest_main
)

include(GoogleTest)
gtest_discover_tests(RunningMaximum)

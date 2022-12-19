# The following lines of boilerplate have to be in your project's
# CMakeLists in this exact order for cmake to work correctly
cmake_minimum_required(VERSION 3.16)

set(WIFI_SSID       "" CACHE STRING "Your Wifi SSID")
set(WIFI_PASS       "" CACHE STRING "Your Wifi password")
set(MQTT_BROKER_URL "" CACHE STRING "Your MQTT broker (mqtt://host:port)")

if (WIFI_SSID STREQUAL "")
	message(FATAL_ERROR "You must set WIFI_SSID")
endif()
if (WIFI_PASS STREQUAL "")
	message(FATAL_ERROR "You must set WIFI_PASS")
endif()
if (MQTT_BROKER_URL STREQUAL "")
	message(FATAL_ERROR "You must set MQTT_BROKER_URL")
endif()

add_compile_definitions(WIFI_SSID="${WIFI_SSID}")
add_compile_definitions(WIFI_PASS="${WIFI_PASS}")
add_compile_definitions(MQTT_BROKER_URL="${MQTT_BROKER_URL}")
include($ENV{IDF_PATH}/tools/cmake/project.cmake)
project(sign)

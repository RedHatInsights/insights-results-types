# insights-results-types

[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

[![GoDoc](https://godoc.org/github.com/RedHatInsights/insights-results-types?status.svg)](https://godoc.org/github.com/RedHatInsights/insights-results-types)
[![GitHub Pages](https://img.shields.io/badge/%20-GitHub%20Pages-informational)](https://redhatinsights.github.io/insights-results-types/)
[![Build Status](https://app.travis-ci.com/RedHatInsights/insights-results-types.svg?branch=master)](https://app.travis-ci.com/RedHatInsights/insights-results-types)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/RedHatInsights/insights-results-types)
[![License](https://img.shields.io/badge/license-Apache-blue)](https://github.com/RedHatInsights/insights-results-types/blob/master/LICENSE)

## Description

Common data types used across the whole CCX data pipeline.

### Class diagram with all types defined in this project

[Class diagram](https://redhatinsights.github.io/insights-results-types/class_diagram.png)

### Note

Please note that this repository should contain just "static" data types, which
means that no runtime code is supposed to be there. Also this repository should
not depend on any other RedHatInsigths packages to avoid circular dependencies.

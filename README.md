<h1 align="center" style="margin-top: 0px;">throttle</h1>

<p align="center">
<a href="https://github.com/space-code/throttle/blob/main/LICENSE"><img alt="Liscence" src="https://img.shields.io/cocoapods/l/service-core.svg?style=flat"></a>
<a href="https://github.com/space-code/throttle"><img alt="CI" src="https://github.com/space-code/throttle/actions/workflows/ci.yml/badge.svg?branch=main"></a>
</p>

## Description
throttle is a handy tool for task throttling.

- [Usage](#usage)
- [Installation](#installation)
- [Communication](#communication)
- [Contributing](#contributing)
- [Author](#author)
- [License](#license)

## Usage

```go
throttler := throttle.New(2 * time.Second)
throttler.Do(func() {
    // your implementation here
})
```

## Installation

```sh
go install github.com/space-code/throttle@latest
```

## Communication
- If you **found a bug**, open an issue.
- If you **have a feature request**, open an issue.
- If you **want to contribute**, submit a pull request.

## Contributing
Bootstrapping development environment

```
make bootstrap
```

Please feel free to help out with this project! If you see something that could be made better or want a new feature, open up an issue or send a Pull Request!

## Author
Nikita Vasilev, nv3212@gmail.com

## License
throttle is available under the MIT license. See the LICENSE file for more info.
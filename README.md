<div id="top"></div>

---

<p align="center">
  <h1 align="center">Gorrola</h1>

  <p align="center">
    A simple Round Robin Load Balancer
  </p>
</p>

<div align="center">
  
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]
  
</div>

<!-- ABOUT THE PROJECT -->
## About The Project

It's a simple load balancer implementing the Round Robin algorithm. This toy project was created because I wanted to 
have something that's built in Go on my GitHub. I use [this article](https://kasvith.me/posts/lets-create-a-simple-lb-go/) as a reference, 
so some parts of the code might look similar, but I assure you I use my own brain most of the time and not just blindly copy what the author wrote.

I have some interesting ideas on how to develop this project further so as to make it my software engineering playground, 
some of which you can find in the [issues page][issues-url]. I don't know if I'm gonna implement those in the future, 
I might though... probably, when I'm bored :p

For now, it has basic functionalities such as:
- Route traffic (I mean, of course)
- Health checks (passive and active)
- Route traffic only to healthy backends
- CLI-based interface

Not very impressive right, but I'm proud of it.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started
### The fast way (Linux only)
1. Download the latest binary from the [release page](https://github.com/danilhendrasr/gorrola/releases)
2. Run it
   ```bash
   chmod +x ./gorrola
   ./gorrola run --backends <your backends URLs here, separated by comma>
   ```
3. Gorrola by defaults will be served at [port 3000](http://localhost:3000)

### Build from source
1. Clone the repo
2. Build the binary
   ```bash
   go build
   ```
3. Run the example backends
   ```bash
   chmod +x scripts/up-backends.sh
   ./scripts/up-backends.sh
   ```
4. Run the binary
   ```bash
   ./gorrola run --backends "http://localhost:8080,http://localhost:8081,http://localhost:8082" -p 3001
   ```
5. Gorrola will be served at [port 3001](http://localhost:3001)

### The docker way? 🫤
I don't provide any docker image at the moment.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See [LICENSE][license-url] for more information.

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[repo-url]: (https://github.com/danilhendrasr/gorrola)
[stars-shield]: https://img.shields.io/github/stars/danilhendrasr/gorrola.svg?style=for-the-badge
[stars-url]: https://github.com/danilhendrasr/gorrola/stargazers
[issues-shield]: https://img.shields.io/github/issues/danilhendrasr/gorrola.svg?style=for-the-badge
[issues-url]: https://github.com/danilhendrasr/gorrola/issues
[license-shield]: https://img.shields.io/github/license/danilhendrasr/gorrola.svg?style=for-the-badge
[license-url]: https://github.com/danilhendrasr/gorrola/blob/main/LICENSE

[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/danilhendrasr/yali4j.svg?style=for-the-badge
[stars-url]: https://github.com/danilhendrasr/yali4j/stargazers
[issues-shield]: https://img.shields.io/github/issues/danilhendrasr/yali4j.svg?style=for-the-badge
[issues-url]: https://github.com/danilhendrasr/yali4j/issues
[license-shield]: https://img.shields.io/github/license/danilhendrasr/yali4j.svg?style=for-the-badge
[license-url]: https://github.com/danilhendrasr/yali4j/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/danilhendrasr
[product-screenshot]: images/screenshot.png

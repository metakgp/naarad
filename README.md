<div id="top"></div>

<!-- PROJECT SHIELDS -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links-->
<div align="center">

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![Wiki][wiki-shield]][wiki-url]

</div>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/metakgp/naarad">
     <img width="140" alt="image" src="https://raw.githubusercontent.com/metakgp/design/main/logos/logo.jpg">
  </a>

  <h3 align="center">Naarad</h3>

  <p align="center">
    <i>Self-hosted ntfy server, delivering notifications to KGPians</i>
    <br />
    <a href="https://naarad.metakgp.org">Website</a>
    ·
    <a href="https://github.com/proffapt/naarad/issues">Request Feature / Report Bug</a>
  </p>
</div>


<!-- TABLE OF CONTENTS -->
<details>
<summary>Table of Contents</summary>

- [About The Project](#about-the-project)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Deployment](#deployment)
- [Contact](#contact)
  - [Maintainer(s)](#maintainers)
  - [creators(s)](#creators)
- [Additional documentation](#additional-documentation)

</details>


<!-- ABOUT THE PROJECT -->
## About The Project
<div align="center">
  <a href="https://github.com/metakgp/naarad">
    <img width="80%" alt="image" src="https://github.com/metakgp/naarad/assets/86282911/21633e13-eb62-4ed6-b188-77514dfdc414">
  </a>
</div>

Narada (Sanskrit: नारद, IAST: Nārada), or Narada Muni, is a sage-divinity, famous in Hindu traditions as a travelling musician and storyteller, who carries news and enlightening wisdom ([source](https://en.wikipedia.org/wiki/Narada)). Our naarad serves news (noticies) to the KGP community. It is a self-hosted [ntfy.sh](https://ntfy.sh) instance with custom configuration. It is protected by [heimdall](https://github.com/metakgp/heimdall), allowing access only to KGPians.<br> 
Following are the features enabled in this ntfy instance (refer [server.yml](./server.yml)):

- Message Caching
- Attachments
- User based authentication
- Keep Alive
- Web Interface
- Webpush notifications
- Logging
- Login/Signup
- iOS Push Notifications
- Rate Limiting

### Custom User Registration

To make the user registration process as seamless as possible we have implemented a custom logic ([frontend](./frontend/) & [backend](./backend/)). The frontend is hosted via github pages with the url - [https://naarad-signup.metakgp.org](https://naarad-signup.metakgp.org), mapped to [https://naarad.metakgp.org/signup](https://naarad.metakgp.org/signup). The logic is as follows:
- User must be authenticated via [heimdall](http://heimdall.metakgp.org)
- Now the magic happens in the backend and registers the user with:
   - `username`: Taken from their institute email
   - `password`: Auto generated
- The credentials are sent to their institute email

> [!Tip]
> To understand the full process of accessing the service, refer [this](./SUBSCRIPTION_INSTRUCTION.md).

<p align="right">(<a href="#top">back to top</a>)</p>

## Getting Started

To set up a local instance of the application, follow the steps below.

### Prerequisites

The following dependencies are required to be installed for the project to function properly:
* [docker](https://docs.docker.com/get-docker/)
* [docker-compose](https://docs.docker.com/compose/install/)

<p align="right">(<a href="#top">back to top</a>)</p>

### Deployment

_Now that the environment has been set up and configured to properly compile and run the project, the next step is to install and configure the project locally on your system._
1. Clone the repository
   ```sh
   git clone https://github.com/metakgp/naarad.git
   ```
2. Copy `.env.example` as `.env` and fill in the required values
3. Build the docker container
   ```sh
   sudo docker compose build
   ```
4. Start the container
   ```sh
   sudo docker compose up -d
   ```
5. The project will be locally deployed at `http://localhost:8000/`

<p align="right">(<a href="#top">back to top</a>)</p>

## Contact

<p>
📫 Metakgp -
<a href="https://slack.metakgp.org">
  <img align="center" alt="Metakgp's slack invite" width="22px" src="https://raw.githubusercontent.com/edent/SuperTinyIcons/master/images/svg/slack.svg" />
</a>
<a href="mailto:metakgp@gmail.com">
  <img align="center" alt="Metakgp's email " width="22px" src="https://raw.githubusercontent.com/edent/SuperTinyIcons/master/images/svg/gmail.svg" />
</a>
<a href="https://www.facebook.com/metakgp">
  <img align="center" alt="metakgp's Facebook" width="22px" src="https://raw.githubusercontent.com/edent/SuperTinyIcons/master/images/svg/facebook.svg" />
</a>
<a href="https://www.linkedin.com/company/metakgp-org/">
  <img align="center" alt="metakgp's LinkedIn" width="22px" src="https://raw.githubusercontent.com/edent/SuperTinyIcons/master/images/svg/linkedin.svg" />
</a>
<a href="https://twitter.com/metakgp">
  <img align="center" alt="metakgp's Twitter " width="22px" src="https://raw.githubusercontent.com/edent/SuperTinyIcons/master/images/svg/twitter.svg" />
</a>
<a href="https://www.instagram.com/metakgp_/">
  <img align="center" alt="metakgp's Instagram" width="22px" src="https://raw.githubusercontent.com/edent/SuperTinyIcons/master/images/svg/instagram.svg" />
</a>
</p>

### Maintainer(s)

The currently active maintainer(s) of this project.

- [Chirag Ghosh](https://github.com/chirag-ghosh)
- [Arpit Bhardwaj](https://github.com/proffapt)

### Creator(s)

Honoring the original creator(s) and ideator(s) of this project.

- [Chirag Ghosh](https://github.com/chirag-ghosh)
- [Arpit Bhardwaj](https://github.com/proffapt)

<p align="right">(<a href="#top">back to top</a>)</p>

## Additional documentation

  - [License](/LICENSE)
  - [Code of Conduct](/.github/CODE_OF_CONDUCT.md)
  - [Security Policy](/.github/SECURITY.md)
  - [Contribution Guidelines](/.github/CONTRIBUTING.md)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->

[contributors-shield]: https://img.shields.io/github/contributors/metakgp/naarad.svg?style=for-the-badge
[contributors-url]: https://github.com/metakgp/naarad/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/metakgp/naarad.svg?style=for-the-badge
[forks-url]: https://github.com/metakgp/naarad/network/members
[stars-shield]: https://img.shields.io/github/stars/metakgp/naarad.svg?style=for-the-badge
[stars-url]: https://github.com/metakgp/naarad/stargazers
[issues-shield]: https://img.shields.io/github/issues/metakgp/naarad.svg?style=for-the-badge
[issues-url]: https://github.com/metakgp/naarad/issues
[license-shield]: https://img.shields.io/github/license/metakgp/naarad.svg?style=for-the-badge
[license-url]: https://github.com/metakgp/naarad/blob/master/LICENSE
[wiki-shield]: https://custom-icon-badges.demolab.com/badge/metakgp_wiki-grey?logo=metakgp_logo&style=for-the-badge
[wiki-url]: https://wiki.metakgp.org
[slack-url]: https://slack.metakgp.org

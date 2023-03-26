<h2 align="center"> Contributing to Road Runner </h2>

We really appreciate your help to develop this project!
Every kind of help is welcome and we want to make contributing as easy as it is possible.

---

Check <b> Table of Contents </b> below to see how to make a particular impact on a project.
All of things pointed out below are descriped in the same order on the bottom of this file.

* Reporting an issue
* Discussing the current state of the code
* Submitting a fix
* Translations
* Proposing new features
* Becoming a maintainer

<br>

<h2 align="center"> Development Process </h2> 

All changes happen through <i> pull requests </i>. Every pull request is welcome, and we invite you to submit your PR
directly <a href="https://github.com/wzslr321/artiver/pulls"> here</a>.
All of pull requests are going to be reviewed and if it goes successfully, these can be merged into the project.

<br>

<h2 align="center"> Quickstart Local Frontend Development </h2>

Do this if you only want to do Flutter related stuff.

* Navigate to the `/client`
* Run `flutter pub get`
* Upgrade flutter version, if it is incompatible with the project
* Run `flutter pub run build_runner build --delete-conflicting-outputs`
* Run `flutter gen-l10n`
* Run `flutter run`

> If you are not sure what is the best way to implement particular thing, consult @wzslr321 or ask on our
<a href="https://discord.gg/Gu2VGcjZfe"> Discord</a>. <b> Please, never feel shy to ask for help! </b>

<br>

<h2 align="center"> Quickstart Local Backend Development </h2>

* Install Docker and docker-compose.

With docker-compose properly configured, you just have to run `docker compose -f docker-compose-dev.yaml up` in the
microservice, e.g. `server/users`.

<br>

<h2 align="center"> Pull Requests </h2>

1. Create an issue describing what you want to do, if it doesn't exist already. Someone will review it shortly, and if
   it is approved,
   you can start working on it and continue to the next steps.
2. Fork the repo and create your own branch, name it according to what it is going to contain.
3. If you have added the code that should be tested, add some test examples or point it out in PR that those are needed.
   Also consider asking for help on
   <a href="https://discord.gg/CSkuazRqKG"> Discord</a>, to have a satisfaction of creating a whole PR by yourself!
4. Ensure to describe your pull request.

> Pull requests trigger checks, make sure it passes all of them.

---

<br>

<h2 align="center"> Reporting an issue </h2> 

You can report the issure directly <a href="https://github.com/wzslr321/ativer/issues/new"> here</a>.

Title of your issue should
contain: `issue(<what is issue related to, e.g. front-end>): Descriptive phrase about the issue`

<i> Issue title example: </i> `issue(state-management): Network connection stream doesn't yield initial value to the state.`

Description should contain as much information as it is possible. Please write also steps to reproduce the issue, so
others can easily see the problem locally on their machines.

<br> 

<h2 align="center"> Discussing the current state of the code </h2>

Every code improvement is extremely welcome and we are always happy to discuss about better solutions. If you think that
some code can be written in a better way, feel free to chat about it on Discord. In case you
are  just curious about any part of code, don't know
why is it done this way, please also ask about it there!

<br> 

<h2 align="center"> Submitting a fix </h2> 

Before submitting a fix, please make sure it has its own issue. If you fixed an issue created by someone else, just
simply use its #id in your pull request title.
Otherwise, before making a pull request, report an issue and describe it there.

Example of fix's PR title: `fix(#63): implement a method to yield initial value to the state.`

<br>

<h2 align="center"> Translations </h2>

You can help to translate the project, even if you don't have any programming knowledge.
All texts used in this project are in `.arb` files located in `/client/lib/l10n/arb`.
To fill a translation gaps in already existing translation file, just edit it.
For example: If you want to add spanish translation, edit the `app_es.arb` file.
If you want to create a translation in new language, simply create a new file in `arb` directory,
with following name: `app_<country_shortcut>.arb`.

English is a primary language of this project, thus all texts that need translation, are
located in `app_en.arb` file.

<b> Translation example: </b>

```bash
/// app_en.arb
{
    "helloWorld": "Hello World!",
    "@helloWorld": {
      "description": "The conventional newborn programmer greeting"
    }
}

/// app_pl.arb
{
    "helloWorld": "Witaj świecie!"
}
```

<h2 align="center"> Proposing a new features </h2>

You can propose a new feature, by simply raising a new issue on GitHub. Make sure to prefix it with `feature:`

<br> 

<h2 align="center"> Becoming a maintainer </h2>

To become a maintainer, you just have to be an active repository contributor or helper. We do appreciate every kind of
help, so if you don't have enough time to write
a code, it would be also awesome if you can help other contributors understand particular code or concepts. Code reviews
are also extremely important part, thus code reviewers
are always welcome.

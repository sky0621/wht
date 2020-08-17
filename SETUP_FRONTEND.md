# Setup Frontend

## env

```
$ yarn -v
1.22.4
$
$ node -v
v12.16.1
$
$ npm -v
6.13.4
```

## create project

https://ja.nuxtjs.org/guides/get-started/installation#using-create-nuxt-app

```
$ yarn create nuxt-app view
yarn create v1.22.4
[1/4] Resolving packages...
[2/4] Fetching packages...
[3/4] Linking dependencies...
[4/4] Building fresh packages...
success Installed "create-nuxt-app@3.2.0" with binaries:
      - create-nuxt-app

create-nuxt-app v3.2.0
✨  Generating Nuxt.js project in view
? Project name: view
? Programming language: TypeScript
? Package manager: Yarn
? UI framework: Vuetify.js
? Nuxt.js modules: (Press <space> to select, <a> to toggle all, <i> to invert selection)
? Linting tools: ESLint, Prettier
? Testing framework: Jest
? Rendering mode: Single Page App
? Deployment target: Static (Static/JAMStack hosting)
? Development tools: (Press <space> to select, <a> to toggle all, <i> to invert selection)
warning nuxt > @nuxt/webpack > @nuxt/babel-preset-app > core-js@2.6.11: core-js@<3 is no longer maintained and not recommended for usage due to the number of issues. Please, upgrade your dependencies to the actu
al version of core-js@3.
warning nuxt > @nuxt/webpack > webpack > watchpack > watchpack-chokidar2 > chokidar@2.1.8: Chokidar 2 will break on node v14+. Upgrade to chokidar 3 with 15x less dependencies.
warning nuxt > @nuxt/webpack > webpack > watchpack > watchpack-chokidar2 > chokidar > fsevents@1.2.13: fsevents 1 will break on node v14+ and could be using insecure binaries. Upgrade to fsevents 2.
warning nuxt > @nuxt/webpack > webpack > micromatch > snapdragon > source-map-resolve > resolve-url@0.2.1: https://github.com/lydell/resolve-url#deprecated
yarn run v1.22.4
$ eslint --ext .js,.vue --ignore-path .gitignore . --fix
Done in 10.71s.

🎉  Successfully created project view

  To get started:

        cd view
        yarn dev

  To build & start for production:

        cd view
        yarn build
        yarn start

  To test:

        cd view
        yarn test


  For TypeScript users. 

  See : https://typescript.nuxtjs.org/cookbook/components/
Done in 186.20s.
```


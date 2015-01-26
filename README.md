# Welcome to Revel Website application in Revel

- The revel www is an application that makes the revel website, manual etc work.
- The revel site in revel is (Dog Food)[http://en.wikipedia.org/wiki/Eating_your_own_dog_food] in software. 


Ideas are to..
    - automate the release cycle and docs
    - create an integrated system
    - make application docs, source code and all work together
    - provide a platform for international

Current revel docs Issues:
- Currently docs are on revel.github.io..
- This restrict things such as no running plugins.. 
- and its written in ruby, and outta out control..in gopher land






Research
=============
This project is a bit of research into
how to still use the markdown and original source
as some of the ideas are good, eg _data/ directory..

It mainly gonna be well dodgy to discover if possible..

But also to be able to mutli-version, and even make the pages dynamic
and make the whole system update itself..

= landing page /
- about/
- download/
- manual/0.13/*
- modules/:source/:ver/*

and stuff like that..

Expect to use..
- bootstrap and go templates
- some golang md parser
- client/server ajax./

- mobile site in jquery mobile..
- other stuff for run, 
- even running examples as proxies, ie revel in revel in revel ;-)

- IRC BOT
- IRC BOT to websockets.
- integration with Github eg auth
- integration with ML
- replace ML with a stackoverflow like Q+A + tags + mini wiki

- any third party modules register and testing





====================
APP REQUIREMENTS:
====================

templates:

make east tags in docs..
 eg 
{{man controllers appconf}}
  = links to controllers and appconf pages in small blue links or something
{{godoc controller config}}

Make alerts.notices easy with
This includes the icon and warning and color configurable..

```alert-important Simple mistake
do not:
- and this and {{man links}}
```


====================
Remove slaving
====================
Ability to check remote git hub..
integrationg with git hub and issues..



    
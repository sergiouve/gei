#### TODOs
- [x] Manage 404 errors when downloading extensions metadata
- [ ] Warning of incompatible versions
- [ ] Allowing to force specific version install
- [ ] Multiple installations on one command
- [ ] Enabling extensions during the installation flow
- [x] Figuring out why the extensions are not detected by gnome
- [ ] Add command to remove extensions
- [ ] Make extensions path configurable by the user
- [x] Add command to list installed extensions
- [ ] ~~Create a middleware to update local database~~
- [x] Implement a dbus channel communication with the shell
- [ ] Allow installation command to use uuid instead of id
- [ ] Abstract all `dbus` calls to a `dbus-client`

---

**_2022-04-18_**

Working on retrieving the gnome version via `dbus` instead of relying on a `gnome-shell --version` cmd call. I got a strange error message which is really tempting me to invest some time investigating how to debug shit in go...

So I successfully used `delve` for debugging but WTF is this:

```go
return shellVersion.Value().(string)
```

So I could just use the `InstallRemoteExtension` from the `dbus` gnome shell API but the user would have to click a confirmation box... At least I think I found the reason I cannot run `global.reexec_self();` through the `Eval` dbus method, I found this in the gnome shell source code:

```js
Eval(code) {
if (!global.context.unsafe_mode)
  return [false, ''];

let returnValue;
let success;
try {
  returnValue = JSON.stringify(eval(code));
  // A hack; DBus doesn't have null/undefined
  if (returnValue == undefined)
    returnValue = '';
      success = true;
     } catch (e) {
       returnValue = `${e}`;
       success = false;
     }
     return [success, returnValue];
}
```

So this approach is not possible without running `unsafe_mode` which is really leaving me with no option to reload the shell...

---

**_2022-04-17_**

Successfully implemented the installed extensions listing via `dbus` yay!

---

**_2022-04-12_**

I finally found the source code for the native `gnome-extensions`. It turns out it is pretty outdated and does not allow to install extensions from the CLOUD TM.

It is also written in `c` and relies in dbus communication with the gnome shell in order to enable/disable/list/and-so-on extensions so it now really looks like I'm not going to be able to avoid a deep dive into bus channels and golang.

https://gitlab.gnome.org/GNOME/gnome-shell/-/blob/main/subprojects/extensions-tool
https://github.com/godbus/dbus

I just found a goldmine
https://gitlab.gnome.org/GNOME/gnome-shell/-/blob/92d3c6e051958b31151bf9538205a71cab6f70d7/data/dbus-interfaces/org.gnome.Shell.Extensions.xml

---

**_2022-04-11_**

Today I'll try to squeeze in some of the features from the aforementioned TODOs lists, completely ignoring the fact that the extensions are still not being enabled. The remove one seems easy engouh...

I'm getting stuck at my system trying to find the go binary since the last session I manually upgraded it to 1.18 and I can't remember which file I have to source.

Turns out its `export PATH=$PATH:/usr/local/go/bin`,

I spent some time doing a bit of folder/packages restructuring and realizing that my initial idea of keeping a local database is going to be needed going forward.

---

**_2022-04-03_**

I have about an hour of free time. I'll try to either make a small refactor or tackle some of the TODOs I annotated last session. I think the 404 one is a good candidate for today.

It looks like I'm diving into golang error handling today.

I just implemented the 404 error handling, just bubbling up the error to the upper most function call. This the correct way? It feels like a cumbersome task in bigger projects.

- [x] Manage 404 errors when downloading extensions metadata
- [ ] Warning of incompatible versions
- [ ] Allowing to force specific version install
- [ ] Multiple installations on one command
- [ ] Enabling extensions during the installation flow
- [ ] Figuring out why the extensions are not detected by gnome
- [ ] Add command to remove extensions

---

**_2022-03-23_**

Apparently I have all the pieces in place, the tool searches, downloads, extracts and deletes the zip file. But GNOME (or rather, GNOME Extensions) does not recognize the installed extension unless I restart the session.

What really need to happen is for the GNOME Shell to restart, so that's what I'll be trying to do today, try to trigger a Shell restart from `gei`. 

Turns out that both `gnome-tweaks` and `gnome-shell-extension-installer` restart the shell via a bus message, I'll have to investigate quite a bit with this one...

---

**_2022-03-20_**

My goal for today is to implement simple functionality like getting the gnome OS version and checking if an extension is already installed.

Ended up creating the helper method AND the actual installation (the shell does not detect them tho, will have to look into it), this last part was surprinsingly easy to do...

Next steps:

- [ ] Manage 404 errors when downloading extensions metadata
- [ ] Warning of incompatible versions
- [ ] Allowing to force specific version install
- [ ] Multiple installations on one command
- [ ] Enabling extensions during the installation flow
- [ ] Figuring out why the extensions are not detected by gnome

---

**_2022-03-18_**

Home alone today, all night ahead of me. It'll be a success if I manage to put in anything more than 30 minutes. I feel a bit lost, I don't remember how I wanted to implement the download.

Got stuck at trying to understand how to share structs (types) across pacakges/files.

---

**_2022-03-14_**

After listening a podcast/interview to the maintainer of the `encoding/json` lib from the standard library I decided to ditch the `jsonparser` package in favor of `encoding/json`. It turns out that it is not that much of a hassle to use it, and it forces you to work with previously defined structs that mirror the json you are expecting to get. I'll have to dedicate some time to the error handling when the http responses either fail or change the json.

---

**_2022-03-02_**

I have about an hour before Diego comes, so I'll try to remember where I left off and maybe make some minor adjustments or refactors.

I ended up refactoring all the client and installer side and getting the client to a single file with a base `fetch` method to wrap the actual http client and exportable `fetches` for the individual http calls. Today felt like I groked golang a little bit better.

---

**_2022-02-23_**

Pretty productive day in terms of awereness gaining I guess. I have the overview of the download process layed out with the Gnome "API". It also felt a bit nicer to read and write code in Go. Not all that much progress code wise though, but I'm getting there, I will have a lot to refactor and FPrize during this learnign process.

---

**_2022-02-21_**

I don't know what I enjoy the most, the actual development or this devlog.

I think this session is going to be mostly about a proper way of formatting the output, doing a quick search on the PCLAiG book I found out about the `Stringer` interface, I tried to copy-paste it's example into my code, which obviosuly did not work. So I'm going for a deep dive into that interface.

I just did a quick read about the `Stringer` interface and ended up doing a minor refactor moving the bussines logic to it's own package (that the correct name?). I'm getting a bit slowed down now by not being able to install `cobra` on the soyuz... 

---

**_2022-02-19_**

First day working on this project using my new tower, soyuz1. I just remembered I left when starting to work on parsing the json response from the extensions API, feels like a mf drag...

Ended up using https://github.com/buger/jsonparser but even then it feels like I'm casting shit on the fly as needed and it really feels like a hack and a total no understanding og the language.

I did manage to print a somewaht formatted results into the terminal so not all that bad for a 40 minutes coding session I guess. I also pushed the code to github (Still thinking about migrating to gitlab and becoming yugarinn) and added a `README` with a xkcd.

I should also read `Powerful Command-Line Applications in Go: Build Fast and Maintainable Tools`.

I'm going to checkout on the cats, haven't seem them for a while...

---

**_2022-02-15_**

Today I'm kind of stuck when importing my own modules from the same project to the `cmd` commands. It seems like I'm doing everything correctly, I checked other projects and my own `go-api-boilerplate` and the setup seems to be the same. Is it possible that `cobra` has rules in place when importing packages into `cmd`?

I have no idea what's going on, the `go.mod` module name is the same as the `import` statement, where am I fucking up?

So it turns out that the package should be the same as the parent folder for the files that define the actual imported functions.

I kinda run out of steam trying to debug that. I'm going to try and push it a bit and see if I can format the output.

Parsing strings to json seems like a bit of a hassle, since the [first article](https://www.sohamkamani.com/golang/json/) I checked recommends to create a struct with the same structre as the json you want to parse, I'm going to leave it here for the day.

---

**_2022-02-14_**

Just started this project, the aim is to provide a POSIX standard command line tool for installing Gnome extensions. I named it `gei`.

This came to me while updating my `dotfiles` project while preparing for my new PC build. I'm also kind of motivated to do it FP style since I read this: [3. Object Oriented Programming, The Trillion Dollar Disaster](3.%20Object%20Oriented%20Programming,%20The%20Trillion%20Dollar%20Disaster.md)

I should really learn what POSIX is actually about tho.

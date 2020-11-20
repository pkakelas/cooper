# cooper

<img src="https://i.ibb.co/NKBzSzS/cooper.png" alt="logo" style="width: 50%" />

Cooper is a fancy program created as an exercise for crawling urls and searching within them.

It consists of:

* A Golang crawler which crawls, parses and stores the crawled data using TF-IDF in an SQLite database.
* A Golang backend which acts as a server for the Cooper search.
* A React.JS frontend which implements the search site called Cooper search.

## Crawling 

Just use the cooper crawler tool
```hs
Welcome to Cooper, an simple and lightweight crawler written in Golang!

           _=,_
        o_/6 /#\
        \__ |##/
        ='|--\
        /   #'-.
        \#|_   _'-. /
         |/ \_( # |"
        C/ ,--___/

Usage: 
  -base_url string 
      The base url where Cooper will start crawling.

  -include_query_params bool
        Should Cooper consider test.com?query and test.com as the same document?
        (default true)

  -limit int
        The maximum sites that Cooper should visit.
        (default 50)

  -load_existed_data
        Whether or not the existing crawled urls should be loaded.
        (default true)

  -server_mode
        Work in server mode for serving data to the cooper frontend.

  -threads int
        How many crawl threads should Cooper use.
        (default 2)
```

---

## Cooper search

* Open the backend with `go run crawler -server-mode`
* Open the frontend with `cd frontend && yarn start`
* If not opened automatically visit: `http://localhost:3000` and search:

<img src="https://i.ibb.co/NTQTpt8/Screenshot-2020-11-20-at-7-10-01-PM.png" alt="logo" style="width: 100%" />
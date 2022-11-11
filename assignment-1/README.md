# Assignment - 1
**PROG2005 - Cloud Technologies**

# Table of contents 
[[_TOC_]]

# Description
This project is about creating a RESTful API using Golang. 
The assignment is to use other APIs and retrieve information in form
of a json body structure. A challenge is also to use these APIs and 
all their features to use these same features on your own API.

The API will use the http://universities.hipolabs.com/ API for
gathering universities and the https://restcountries.com/ API for
gathering countries and their details.

The created API is being deployed to the Heroku service on this 
link: **URL**

The API also contains a diagnostics overview checking if the other
APIs have a status code of 200 (OK) and return this. It is formatted
as a json file and will also show the programs uptime.

# Cloning
To use this software you will have to clone this from GitHub. You will need to enter a terminal and enter the following:

    1. Cloning the Remote Repo to your Local host. example: git clone https://github.com/user-name/repository.git.
    2. Pulling the Remote Repo to your Local host. First you have to create a git local repo by, example: git init or git init repo-name then, git pull https://github.com/user-name/repository.git.


# Useage

To use this API you will need three URLs.

1. /unisearcher/v1/uniinfo/
2. /unisearcher/v1/neighbourunis/
3. /unisearcher/v1/diag/

To use these you will need to enter the Heroku URL followed
by one of the extensions.

## 1 /unisearcher/v1/uniinfo/. 
By using this API you will
be able to search through the hippoLabs API and combining 
the universities gathered with its country, using the partial or
full name of the university. You who use the 
API will then retrieve a json splice of **all** universities 
matching your partial or full name.

It will look like this:

    Method: GET
    Path: uniinfo/{:partial_or_complete_university_name}/

Where you, the user will use the **GET** command by using the
path and setting the partial og complete university name.

Body (Example):
```
[
  {
      "name": "Norwegian University of Science and Technology", 
      "country": "Norway",
      "isocode": "NO",
      "webpages": ["http://www.ntnu.no/"],
      "languages": {"nno": "Norwegian Nynorsk",
                    "nob": "Norwegian Bokmål",
                    "smi": "Sami"},
      "map": "https://www.openstreetmap.org/relation/2978650"
  },
  ...
]
```


## 2 /unisearcher/v1/neighbourunis/. 
In this API you will be able to get the same universities as in 
method 1. but you'll also get the neighbouring countries' 
universities which also contains a partial or complete university
name.

It will look like this:

    Method: GET
    Path: neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}

Where you, the user will use the **GET** command by using the
path and setting the country to get the borders from by entering the
partial og complete university name. You can also optionally limit
to _x_ amount of universities shown.

Body (Example):
```
[
  {
      "name": "Norwegian University of Science and Technology", 
      "country": "Norway",
      "isocode": "NO",
      "webpages": ["http://www.ntnu.no/"],
      "languages": {"nno": "Norwegian Nynorsk",
                    "nob": "Norwegian Bokmål",
                    "smi": "Sami"},
      "map": "https://www.openstreetmap.org/relation/2978650"
  },
  {
      "name": "Swedish University of Agricultural Sciences", 
      "country": "Sweden",
      "isocode": "SE",
      "webpages": ["http://www.slu.se/"],
      "languages": {"swe":"Swedish"},
      "map": "https://www.openstreetmap.org/relation/52822"
  },
  ...
]
```
## 3 /unisearcher/v1/diag/. 

The API also contains a diagnostics handler which can be used for
checking the APIs we are gathering information from. And also
seeing the uptime of the API you're hosting.

    Method: GET
    Path: diag/

Where you, the user will use the **GET** command by using the
path.

Body:
```
{
   "universitiesapi": "<http status code for universities API>",
   "countriesapi": "<http status code for restcountries API>",
   "version": "v1",
   "uptime": <time in seconds from the last service restart>
}
```

# Known issues
A known issue is when typing in for example China in the 
neighbouringunis API, Taiwan won't show up, unless this is an issue
with the other API, this is a fault of mine.

The uniinfo API is a bit slow. If you search something like 
"university", it will use a long time to find the results. It
will find the correct results and return these, it's just slow.

# Heroku deployment
The web service is hosted on Heroku. 
It is only necessary to use the following URL and write 
the valid endpoint(s) and values:
URL to the deployed Heroku service:
https://assignment-1-heroku.herokuapp.com/
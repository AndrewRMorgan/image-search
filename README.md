# FreeCodeCamp API Basejump: Image Search Abstraction Layer
## User stories:
1. I can get the image URLs, alt text and page urls for a set of images relating to a given search string.
2. I can paginate through the responses by adding a ?offset=2 parameter to the URL.
3. I can get a list of the most recently submitted search strings.

## Image Search Example:

`https://mysterious-plateau-36613.herokuapp.com/api/imagesearch/funny%20dog%20photos`

## Output:

```js
{ url: "https://i.ytimg.com/vi/GF60Iuh643I/hqdefault.jpg",
  snippet: "Funny Dogs - A Funny Dog Videos Compilation 2015 - YouTube",
  thumbnail: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ6qlPomJo0DJyIx5gyvyeqQZyRJHkcPdmnAXXVwRGN_Z6v63XoHQvcj3w",
  context: "https://www.youtube.com/watch?v=GF60Iuh643I" }
```

## Recently Searched Strings Example::

`https://mysterious-plateau-36613.herokuapp.com/api/latest/imagesearch`

### Output:

```js
{ term: "funny dog photos", when: "2017-05-30T10:01:12Z" }
```

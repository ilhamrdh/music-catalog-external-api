package spotify

var mockResponse = `
{
  "tracks": {
    "href": "https://api.spotify.com/v1/search?offset=0&limit=10&query=bernadya&type=track&market=ID&locale=id-ID,id;q%3D0.9,en-US;q%3D0.8,en;q%3D0.7,de;q%3D0.6",
    "limit": 10,
    "next": "https://api.spotify.com/v1/search?offset=10&limit=10&query=bernadya&type=track&market=ID&locale=id-ID,id;q%3D0.9,en-US;q%3D0.8,en;q%3D0.7,de;q%3D0.6",
    "offset": 0,
    "previous": null,
    "total": 818,
    "items": [
      {
        "album": {
          "album_type": "single",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/47z98pKd71yIbgXwe9LPVC"
              },
              "href": "https://api.spotify.com/v1/artists/47z98pKd71yIbgXwe9LPVC",
              "id": "47z98pKd71yIbgXwe9LPVC",
              "name": "Bernadya",
              "type": "artist",
              "uri": "spotify:artist:47z98pKd71yIbgXwe9LPVC"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/5K8xqV7MCe3UIfedVAlZSe"
          },
          "href": "https://api.spotify.com/v1/albums/5K8xqV7MCe3UIfedVAlZSe",
          "id": "5K8xqV7MCe3UIfedVAlZSe",
          "images": [
            {
              "height": 640,
              "width": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b273cb1e5f7d0942bf9196c1e711"
            },
            {
              "height": 300,
              "width": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e02cb1e5f7d0942bf9196c1e711"
            },
            {
              "height": 64,
              "width": 64,
              "url": "https://i.scdn.co/image/ab67616d00004851cb1e5f7d0942bf9196c1e711"
            }
          ],
          "is_playable": true,
          "name": "Terlintas",
          "release_date": "2023-06-22",
          "release_date_precision": "day",
          "total_tracks": 5,
          "type": "album",
          "uri": "spotify:album:5K8xqV7MCe3UIfedVAlZSe"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/47z98pKd71yIbgXwe9LPVC"
            },
            "href": "https://api.spotify.com/v1/artists/47z98pKd71yIbgXwe9LPVC",
            "id": "47z98pKd71yIbgXwe9LPVC",
            "name": "Bernadya",
            "type": "artist",
            "uri": "spotify:artist:47z98pKd71yIbgXwe9LPVC"
          }
        ],
        "disc_number": 1,
        "duration_ms": 200476,
        "explicit": false,
        "external_ids": {
          "isrc": "IDA682300214"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/7zOVh5fGpEwSbZd0g9z80B"
        },
        "href": "https://api.spotify.com/v1/tracks/7zOVh5fGpEwSbZd0g9z80B",
        "id": "7zOVh5fGpEwSbZd0g9z80B",
        "is_local": false,
        "is_playable": true,
        "name": "Satu Bulan",
        "popularity": 83,
        "preview_url": null,
        "track_number": 5,
        "type": "track",
        "uri": "spotify:track:7zOVh5fGpEwSbZd0g9z80B"
      },
      {
        "album": {
          "album_type": "album",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/47z98pKd71yIbgXwe9LPVC"
              },
              "href": "https://api.spotify.com/v1/artists/47z98pKd71yIbgXwe9LPVC",
              "id": "47z98pKd71yIbgXwe9LPVC",
              "name": "Bernadya",
              "type": "artist",
              "uri": "spotify:artist:47z98pKd71yIbgXwe9LPVC"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/5agEAxt8vks5Xk0NfEbI5D"
          },
          "href": "https://api.spotify.com/v1/albums/5agEAxt8vks5Xk0NfEbI5D",
          "id": "5agEAxt8vks5Xk0NfEbI5D",
          "images": [
            {
              "height": 640,
              "width": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b27327693aaf059002bba3a7655a"
            },
            {
              "height": 300,
              "width": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e0227693aaf059002bba3a7655a"
            },
            {
              "height": 64,
              "width": 64,
              "url": "https://i.scdn.co/image/ab67616d0000485127693aaf059002bba3a7655a"
            }
          ],
          "is_playable": true,
          "name": "Sialnya, Hidup Harus Tetap Berjalan",
          "release_date": "2024-06-23",
          "release_date_precision": "day",
          "total_tracks": 8,
          "type": "album",
          "uri": "spotify:album:5agEAxt8vks5Xk0NfEbI5D"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/47z98pKd71yIbgXwe9LPVC"
            },
            "href": "https://api.spotify.com/v1/artists/47z98pKd71yIbgXwe9LPVC",
            "id": "47z98pKd71yIbgXwe9LPVC",
            "name": "Bernadya",
            "type": "artist",
            "uri": "spotify:artist:47z98pKd71yIbgXwe9LPVC"
          }
        ],
        "disc_number": 1,
        "duration_ms": 182903,
        "explicit": false,
        "external_ids": {
          "isrc": "IDA682400233"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/2gcMYiZzzmzoF8PPAfL3IO"
        },
        "href": "https://api.spotify.com/v1/tracks/2gcMYiZzzmzoF8PPAfL3IO",
        "id": "2gcMYiZzzmzoF8PPAfL3IO",
        "is_local": false,
        "is_playable": true,
        "name": "Untungnya, Hidup Harus Tetap Berjalan",
        "popularity": 82,
        "preview_url": null,
        "track_number": 8,
        "type": "track",
        "uri": "spotify:track:2gcMYiZzzmzoF8PPAfL3IO"
      }
    ]
  }
}
`

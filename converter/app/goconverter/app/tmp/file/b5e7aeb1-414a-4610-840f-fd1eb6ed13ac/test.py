import requests
import json
import threading
import time

class Artist:
    def __init__(self, artist):
        self.artist = artist

class Discorgs:
    def __init__(self):
        self.lock = threading.Lock()
        self.artist_list = []
        self.BASE_URL = "https://api.discogs.com"
        self.USER_AGENT = "YOUR_APP_NAME"
        self.TOKEN = "bwvuoSQgNgGeFAeDYUjxOnHcNbCdbdbpErkhdVvb"

        self.headers = {"User-Agent": self.USER_AGENT, "Authorization": f"Discogs token={self.TOKEN}"}
        response = requests.get("https://api.discogs.com/", headers=self.headers)
        python_object = json.loads(response.text)
        self.artists_nbr = int(python_object['statistics']["artists"])

    def background_lock(self):
        with self.lock:
            time.sleep(60)

    def get_data_from_url(self, url):
        with self.lock:
            response = requests.get(url, headers=self.headers)
            rate_limit_remaining = response.headers.get("X-Discogs-Ratelimit-Remaining")
        if int(rate_limit_remaining) < 5:
            background_thread = threading.Thread(target=self.background_lock)
            background_thread.start()

        if response.status_code == 200:
                return response.json()
        return None
    
    def get_artist(self, artist_id):
        return self.get_data_from_url("https://api.discogs.com/artists/"+ str(artist_id) + "/releases")

    def calculate_pagination(self, data):
        e = data["pagination"]
        return int(e["pages"]) * int(e["per_page"])
    
    def to_json_file(self):
        json_data = json.dumps(self.artist_list)
        with open('people.json', 'w') as json_file:
            json_file.write(json_data)

def maines01():
    diso = Discorgs()
    for i in range(1, 10):
        data = diso.get_artist(i)
        if data == None:
            continue
        per_page = diso.calculate_pagination(data)
        artist_content = diso.get_data_from_url("https://api.discogs.com/artists/" + str(i) + "/releases?page=1&per_page=" + str(per_page))
        diso.artist_list.append(Artist(artist_content))

if __name__ == "__main__":
    main()
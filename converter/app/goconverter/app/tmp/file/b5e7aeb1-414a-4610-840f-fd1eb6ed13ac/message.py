import os
import json
import argparse
import unicodedata
import time
from tqdm import tqdm
import yt_dlp
import requests

BASE_URL = "https://api.discogs.com"
USER_AGENT = "YOUR_APP_NAME"
TOKEN = "bwvuoSQgNgGeFAeDYUjxOnHcNbCdbdbpErkhdVvb"

headers = {"User-Agent": USER_AGENT, "Authorization": f"Discogs token={TOKEN}"}


def get_data_from_url(url):
    response = requests.get(url, headers=headers)
    rate_limit_remaining = response.headers.get("X-Discogs-Ratelimit-Remaining")

    print("RATE : " + rate_limit_remaining)

    if int(rate_limit_remaining) < 5:
        print("Approaching rate limit. Pausing for 60 seconds...")
        time.sleep(60)

    if response.status_code == 200:
        return response.json()
    return None


def save_to_json(data, filename):
    try:
        with open(filename, "w") as file:
            json.dump(data, file, indent=4)
    except Exception as e:
        print(e)


def download_video(video_url, artist_name, album_name, error_log):
    options = {
        "format": "bestaudio/best",
        "postprocessors": [
            {
                "key": "FFmpegExtractAudio",
                "preferredcodec": "mp3",
                "preferredquality": "192",
            }
        ],
        "outtmpl": os.path.join(
            os.getcwd(), artist_name, album_name, "%(title)s.%(ext)s"
        ),
        "embed-thumbnail": True,
    }

    with yt_dlp.YoutubeDL(options) as ydl:
        try:
            ydl.download([video_url])
        except yt_dlp.utils.DownloadError as e:
            error_message = (
                f"Failed to download video for {artist_name} - {album_name}: {str(e)}"
            )
            print(error_message)
            error_log.write(error_message + "\n")


def download_thumbnail(thumbnail_url, artist_name):
    if thumbnail_url:
        options = {"outtmpl": os.path.join(os.getcwd(), artist_name, "thumbnail.jpg")}

        with yt_dlp.YoutubeDL(options) as ydl:
            try:
                ydl.download([thumbnail_url])
            except yt_dlp.utils.DownloadError as e:
                error_message = (
                    f"Failed to download thumbnail for {artist_name}: {str(e)}"
                )
                print(error_message)


def main():
    # Créer un analyseur d'arguments
    parser = argparse.ArgumentParser(
        description="Scrape Discogs artist data and optionally download videos."
    )

    # Ajouter un argument pour le téléchargement des vidéos
    parser.add_argument(
        "--download-videos", action="store_true", help="Download videos for artists"
    )

    # Obtenir les arguments de la ligne de commande
    args = parser.parse_args()

    page_number = 1
    all_artists = []

    while page_number <= 100:  
        artists_url = (
            f"{BASE_URL}/database/search?type=artist&per_page=100&page={page_number}"
        )
        artists_data = get_data_from_url(artists_url)

        if not artists_data:
            break  

        artists_on_current_page = artists_data.get("results", [])
        all_artists.extend(artists_on_current_page)
        page_number += 1

        # Parcourir tous les artistes
        for artist_data in tqdm(all_artists, desc="Processing artists"):
            artist_name = artist_data["title"]

            # Normalisez le nom de l'artiste pour éviter les caractères spéciaux dans les noms de dossier
            artist_name_normalized = (
                unicodedata.normalize("NFKD", artist_name)
                .encode("ascii", "ignore")
                .decode("utf-8")
            )
            artist_name_normalized = artist_name_normalized.replace(" ", "_")

            # Vérifiez si le dossier de l'artiste existe déjà, et si c'est le cas, sautez cet artiste
            artist_folder = os.path.join(os.getcwd(), artist_name_normalized)
            if os.path.exists(artist_folder):
                print(
                    f"Le dossier pour l'artiste '{artist_name}' existe déjà. Sauter cet artiste."
                )
                continue

            # Créer un dossier pour l'artiste
            os.makedirs(artist_folder, exist_ok=True)

            # Créer un fichier JSON contenant les informations de l'artiste à l'intérieur du dossier de l'artiste
            artist_info_filename = os.path.join(artist_folder, "artist_info.json")
            save_to_json(artist_data, artist_info_filename)

            # Télécharger la thumbnail de l'artiste
            download_thumbnail(artist_data.get("thumb", ""), artist_name_normalized)

            releases_url = artist_data["resource_url"] + "/releases"
            release_data = get_data_from_url(releases_url)
            if release_data is not None:
                releases = release_data.get("releases", [])
            else:
                print("Échec de la récupération des données de sortie.")
                continue  # Passer à l'artiste suivant en cas d'erreur

            for release in tqdm(
                releases, desc=f"Traitement des albums pour {artist_name}"
            ):
                detailed_release_data = get_data_from_url(release["resource_url"])

                album_name = release["title"][:100]

                invalid_folder_characters = str.maketrans(
                    {
                        " ": "_",
                        ":": "_",
                        "?": "_",
                        "*": "_",
                        "/": "_",
                        "<": "_",
                        ">": "_",
                        "|": "_",
                        '"': "_",
                        "\\": "_",
                        "'": "_",
                        "\t": "_",
                    }
                )

                album_name_normalized = (
                    unicodedata.normalize("NFKD", album_name)
                    .encode("ascii", "ignore")
                    .decode("utf-8")
                )

                album_name_normalized = album_name_normalized.translate(
                    invalid_folder_characters
                )

                # Create the album folder with a sanitized name
                album_folder = os.path.join(artist_folder, album_name_normalized)
                os.makedirs(album_folder, exist_ok=True)

                album_info_filename = os.path.join(album_folder, "album_info.json")
                album_data = {
                    "Album_Name": album_name,
                    "Titles": [detailed_release_data],
                }
                save_to_json(album_data, album_info_filename)

                # Vérifiez s'il y a des vidéos et téléchargez-les si nécessaire
                if args.download_videos and "videos" in detailed_release_data:
                    for video in detailed_release_data["videos"]:
                        video_url = video.get("uri")
                        if video_url:
                            download_video(
                                video_url,
                                artist_name_normalized,
                                album_name_normalized,
                                error_log,
                            )


if __name__ == "__main__":
    main()

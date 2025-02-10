### `DOCUMENTATION.md`

```markdown
# Documentation: Music Playlist Generator

This documentation outlines the technical details of the Music Playlist Generator project.

---

## 1. Overview
The application provides users with an easy way to create and manage music playlists. It integrates with the Spotify API for fetching trending songs and the Lyrics.ovh API for lyrics.

---

## 2. Backend

### Routes
- **GET `/api/trending`**  
  - **Description**: Fetches a list of trending songs from Spotify.
  - **Response**: JSON array of songs.
- **GET `/api/playlists`**  
  - **Description**: Retrieves saved playlists.
  - **Response**: JSON array of playlists.
- **POST `/api/save`**  
  - **Description**: Saves a playlist.
  - **Request Body**:
    ```json
    {
      "name": "Playlist Name",
      "songs": [
        {"title": "Song Title", "artist": "Artist Name", "album": "Album Name"}
      ]
    }
    ```

---

## 3. Middleware
- **LoggingMiddleware**: Logs HTTP method, request URI, and response time.
- **JSONContentTypeMiddleware**: Ensures JSON `Content-Type` for `/api` routes.

---

## 4. Frontend
- **HTML**: The structure of the app.
- **CSS**: Styled with responsive design for modern devices.
- **JavaScript**: Handles API calls and dynamic UI updates.

---

## 5. Services
- **Spotify Integration**:
  - Uses `FetchTrendingSongs` to retrieve Spotify playlist tracks.
  - Requires Spotify Client ID and Secret.
- **Lyrics Integration**:
  - Fetches lyrics for a specific song using `FetchLyrics`.

---

## 6. Environment Variables
Create a `.env` file:
```plaintext
SPOTIFY_CLIENT_ID=<your_client_id>
SPOTIFY_CLIENT_SECRET=<your_client_secret>
REDIRECT_URL=http://localhost:3090

---

7. Error Handling
Invalid Request: Returns a 400 Bad Request with an error message.
API Errors: Returns 500 Internal Server Error for third-party failures.

---

8. Future Improvements
Add user authentication.
Implement persistent storage (e.g., PostgreSQL).
Allow users to share playlists.

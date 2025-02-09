// Global variables for state management
let isLoading = false;

// Utility function for showing loading state
function setLoading(loading) {
    isLoading = loading;
    const buttons = document.querySelectorAll('button');
    buttons.forEach(button => {
        button.disabled = loading;
    });
}

// Utility function for showing notifications
function showNotification(message, isError = false) {
    const alertDiv = document.createElement('div');
    alertDiv.className = `notification ${isError ? 'error' : 'success'}`;
    alertDiv.style.cssText = `
        position: fixed;
        top: 20px;
        right: 20px;
        padding: 15px 25px;
        border-radius: 5px;
        color: white;
        background-color: ${isError ? '#ff4444' : '#44b544'};
        z-index: 1000;
        transition: opacity 0.3s ease;
    `;
    alertDiv.textContent = message;
    document.body.appendChild(alertDiv);

    // Remove notification after 3 seconds
    setTimeout(() => {
        alertDiv.style.opacity = '0';
        setTimeout(() => {
            document.body.removeChild(alertDiv);
        }, 300);
    }, 3000);
}

// Fetch trending songs from the API
async function fetchTrendingSongs() {
    try {
        setLoading(true);
        const response = await fetch('/api/trending');
        if (!response.ok) {
            throw new Error('Failed to fetch trending songs');
        }
        const songs = await response.json();
        
        const trendingList = document.getElementById('trending-list');
        trendingList.innerHTML = ''; // Clear existing list
        
        songs.forEach(song => {
            const li = document.createElement('li');
            const artistName = song.Artists && song.Artists[0] ? song.Artists[0].Name : 'Unknown Artist';
            li.textContent = `${song.Name} - ${artistName}`;
            
            // Add play button or other controls if needed
            const controlsDiv = document.createElement('div');
            controlsDiv.className = 'song-controls';
            // Add controls here if needed
            
            li.appendChild(controlsDiv);
            trendingList.appendChild(li);
        });
    } catch (error) {
        console.error('Error fetching trending songs:', error);
        showNotification('Failed to load trending songs', true);
    } finally {
        setLoading(false);
    }
}

// Fetch saved playlists from the API
async function fetchPlaylists() {
    try {
        setLoading(true);
        const response = await fetch('/api/playlists');
        if (!response.ok) {
            throw new Error('Failed to fetch playlists');
        }
        const playlists = await response.json();
        
        const playlistList = document.getElementById('playlist-list');
        playlistList.innerHTML = ''; // Clear existing list
        
        playlists.forEach(playlist => {
            const li = document.createElement('li');
            li.textContent = playlist.name || playlist; // Handle both object and string responses
            
            // Add delete button
            const deleteButton = document.createElement('button');
            deleteButton.textContent = '×';
            deleteButton.className = 'delete-playlist';
            deleteButton.onclick = () => deletePlaylist(playlist.id || playlist);
            
            li.appendChild(deleteButton);
            playlistList.appendChild(li);
        });
    } catch (error) {
        console.error('Error fetching playlists:', error);
        showNotification('Failed to load playlists', true);
    } finally {
        setLoading(false);
    }
}

// Save a new playlist
async function savePlaylist(name) {
    try {
        setLoading(true);
        const response = await fetch('/api/save', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name: name }),
        });
        
        if (!response.ok) {
            throw new Error('Failed to save playlist');
        }
        
        const result = await response.json();
        showNotification(result.message || 'Playlist saved successfully!');
        await fetchPlaylists(); // Refresh the playlist list
    } catch (error) {
        console.error('Error saving playlist:', error);
        showNotification('Failed to save playlist', true);
    } finally {
        setLoading(false);
    }
}

// Delete a playlist (if implemented on backend)
async function deletePlaylist(playlistId) {
    if (!confirm('Are you sure you want to delete this playlist?')) {
        return;
    }
    
    try {
        setLoading(true);
        const response = await fetch(`/api/playlists/${playlistId}`, {
            method: 'DELETE',
        });
        
        if (!response.ok) {
            throw new Error('Failed to delete playlist');
        }
        
        showNotification('Playlist deleted successfully!');
        await fetchPlaylists(); // Refresh the playlist list
    } catch (error) {
        console.error('Error deleting playlist:', error);
        showNotification('Failed to delete playlist', true);
    } finally {
        setLoading(false);
    }
}

// Event Listeners
document.addEventListener('DOMContentLoaded', () => {
    // Load initial data
    fetchTrendingSongs();
    fetchPlaylists();
    
    // Setup form submission
    const playlistForm = document.getElementById('playlist-form');
    if (playlistForm) {
        playlistForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const nameInput = document.getElementById('playlist-name');
            const playlistName = nameInput.value.trim();
            
            if (!playlistName) {
                showNotification('Please enter a playlist name', true);
                return;
            }
            
            await savePlaylist(playlistName);
            nameInput.value = ''; // Clear the input
        });
    }
    
    // Optional: Add refresh buttons
    const refreshButtons = document.querySelectorAll('.refresh-button');
    refreshButtons.forEach(button => {
        button.addEventListener('click', () => {
            if (button.dataset.target === 'trending') {
                fetchTrendingSongs();
            } else if (button.dataset.target === 'playlists') {
                fetchPlaylists();
            }
        });
    });
});

// Optional: Add keyboard shortcuts
document.addEventListener('keydown', (e) => {
    // Ctrl/Cmd + R to refresh both lists
    if ((e.ctrlKey || e.metaKey) && e.key === 'r') {
        e.preventDefault();
        fetchTrendingSongs();
        fetchPlaylists();
    }
});
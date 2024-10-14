<template>
  <div class="container">
    <!-- Authentication card when the user is not authenticated -->
    <div class="auth-card" v-if="!isAuthenticated">
      <h1 class="title">Bem vindo!</h1>
      <p>Para continuar é necessário fazer login.</p>
      <button class="btn primary" @click="login">Login</button>
    </div>

    <!-- Display user info and actions when authenticated -->
    <div class="auth-card" v-if="isAuthenticated">
      <h1 class="title">Olá, {{ user.profile.given_name }}!</h1>
      <p class="email">{{ user.profile.email }}</p>
      <p style="color: green; font-weight: bold">Você esta autenticado.</p>
      <div class="actions">
        <button class="btn logout" @click="logout">Logout</button>
        <button class="btn secondary" @click="getPersonList">Ver lista de pessoas</button>
      </div>

      <!-- Person list displayed when data is available -->
      <div v-if="personList.length > 0" class="person-list">
        <h2>Pessoas:</h2>
        <div v-for="p in personList" :key="p" class="person-card">
          <div class="person-info">
            <label>First Name:</label>
            <span>{{ p.first_name }}</span>
          </div>
          <div class="person-info">
            <label>Last Name:</label>
            <span>{{ p.last_name }}</span>
          </div>
          <div class="person-info">
            <label>Age:</label>
            <span>{{ p.age }}</span>
          </div>
          <div class="person-info">
            <label>Document:</label>
            <span>{{ p.document }}</span>
          </div>
          <div class="person-info">
            <label>Address:</label>
            <span>{{ p.address }}</span>
          </div>
          <div class="person-info">
            <label>Phone:</label>
            <span>{{ p.phone }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { useOidc } from './services/oidc'
import { onMounted, ref } from 'vue'
import axios from 'axios'

export default {
  setup() {
    // Destructuring the OIDC hook to manage authentication
    const { user, isAuthenticated, login, logout, handleCallback } = useOidc()

    // Reactive reference to hold the list of people
    const personList = ref([])

    // Runs when the component is mounted (similar to componentDidMount in React)
    onMounted(() => {
      const url = window.location.href
      if (url.includes('callback')) {
        // Handle OIDC callback if the URL includes 'callback'
        handleCallback()
      }
    })

    // Fetches the list of people from the API
    const getPersonList = async () => {
      // Check if the user is authenticated and has an access token
      if (!user.value || !user.value.access_token) {
        console.error('User not authenticated or token is unavailable')
        return
      }

      try {
        // API request to fetch people data
        const { data } = await axios.get('https://dev.fibralink.net.br/api/people', {
          // Attaching access token to the request
          headers: {
            Authorization: `Bearer ${user.value.access_token}`
          }
        })

        // Setting the retrieved data to the reactive `personList`
        personList.value = data
        console.log('Person list:', personList.value)
      } catch (error) {
        // Error handling in case the request fails
        console.error('Error fetching person list:', error)
      }
    }

    // Returning variables and methods to be used in the template
    return {
      user,
      isAuthenticated,
      login,
      logout,
      getPersonList,
      personList
    }
  }
}
</script>

<style scoped>
html,
body {
  height: 100%;
  padding: 0;
  margin: 0;
}

/* Container styling to center content vertically and horizontally */
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100%;
}

/* Authentication card styles */
.auth-card {
  background-color: #ffffff;
  padding: 100px;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 600px;
}

/* Title and paragraph styles */
h1 {
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
}

h2 {
  color: #333;
}

p {
  font-size: 16px;
  color: #666;
  margin-bottom: 30px;
}

.email {
  font-size: 14px;
  color: #666;
  margin-bottom: 20px;
}

/* Button styles */
.btn {
  background-color: #007bff;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

/* Button hover effect */
.btn:hover {
  background-color: #0056b3;
}

.actions {
  display: flex;
  gap: 10px;
  justify-content: center;
  margin-bottom: 20px;
}

/* Person list styles */
.person-list {
  margin-top: 20px;
  text-align: left;
}

.person-card {
  background-color: #f9f9f9;
  padding: 10px;
  margin-bottom: 10px;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.person-info {
  display: flex;
  justify-content: space-between;
  margin-bottom: 5px;
}

/* Label and text styles for person details */
label {
  font-weight: bold;
  color: #333;
}

span {
  color: #555;
}

/* Button styles for primary, secondary, and logout actions */
.btn.primary {
  background-color: #007bff;
}

.btn.secondary {
  background-color: #6c757d;
}

.btn.logout {
  background-color: #dc3545;
}
</style>

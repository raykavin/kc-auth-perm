<template>
  <div class="callback-container">
    <div class="loading-message" v-if="!error">
      <p>Processando o login...</p>
    </div>
    <div class="error-message" v-if="error">
      <h2>Erro durante a autenticação</h2>
      <p>{{ error }}</p>
      <button @click="goHome">Voltar para a página inicial</button>
    </div>
  </div>
</template>

<script>
import { useOidc } from './services/oidc'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

export default {
  name: 'AuthCallback',
  setup() {
    const error = ref(null)
    const { handleCallback } = useOidc()
    const router = useRouter()

    const goHome = () => {
      router.push('/')
    }

    onMounted(async () => {
      try {
        setTimeout(async () => {
          await handleCallback()
          goHome()
        }, 3000)
      } catch (err) {
        console.error('Erro ao processar o callback do OIDC:', err)
        error.value = err.message || 'Ocorreu um erro desconhecido'
      }
    })

    return {
      error,
      goHome
    }
  }
}
</script>

<style scoped>
.callback-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f0f4f8;
}

.loading-message,
.error-message {
  font-size: 18px;
  color: #333;
  padding: 20px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.error-message {
  text-align: center;
}

.error-message h2 {
  color: #dc3545;
}

button {
  margin-top: 20px;
  background-color: #007bff;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

button:hover {
  background-color: #0056b3;
}
</style>

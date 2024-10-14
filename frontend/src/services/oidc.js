import { ref, onMounted } from 'vue'
import { UserManager } from 'oidc-client-ts'

export function useOidc() {
  const user = ref(null)
  const isAuthenticated = ref(false)

  // OIDC configuration using environment variables
  const oidcConfig = {
    authority: import.meta.env.VITE_SSO_AUTHORITY_ADDRESS,
    client_id: import.meta.env.VITE_SSO_AUTHORITY_CLIENT_ID,
    response_type: import.meta.env.VITE_SSO_AUTHORITY_RESPONSE_TYPE,
    scope: import.meta.env.VITE_SSO_AUTHORITY_SCOPE,
    redirect_uri: import.meta.env.VITE_CALLBACK_URI,
    post_logout_redirect_uri: import.meta.env.VITE_POST_LOGOUT_REDIRECT_URI,
    revokeTokenTypes: ['refresh_token'],
    automaticSilentRenew: false
  }

  const userManager = new UserManager(oidcConfig)

  // Function to handle login via OIDC redirect
  const login = async () => {
    try {
      await userManager.signinRedirect()
    } catch (error) {
      console.error('Erro durante o login:', error)
    }
  }

  // Function to handle logout via OIDC redirect
  const logout = async () => {
    try {
      await userManager.signoutRedirect()
    } catch (error) {
      console.error('Erro ao relizar o logout:', error)
    }
  }

  // Function to handle OIDC callback after redirection
  const handleCallback = async () => {
    try {
      const userResponse = await userManager.signinRedirectCallback()
      user.value = userResponse
      isAuthenticated.value = !!userResponse
    } catch (error) {
      console.error('Erro ao manipular retorno da chamada OIDC:', error)
      throw error // Re-throw the error if needed for external handling
    }
  }

  // Function to fetch the current user when the component is mounted
  const fetchUser = async () => {
    try {
      const userResponse = await userManager.getUser()
      if (userResponse) {
        user.value = userResponse
        isAuthenticated.value = !!userResponse
      }
    } catch (error) {
      console.error('Erro ao obter informações do usuário logado:', error)
    }
  }

  // Fetch the user when the component is mounted
  onMounted(fetchUser)

  return {
    user,
    isAuthenticated,
    login,
    logout,
    handleCallback
  }
}

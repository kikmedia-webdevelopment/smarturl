const devURL = 'http://localhost:8080'

export default {
    // probs dev if not present
    baseUrl: window.config?.apiurl ?? devURL,
    apiUrl: window.config?.apiurl ? `${window.config.apiurl}/api` : `${devURL}/api`
}
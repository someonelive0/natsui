import { createApp } from 'vue'

// Vuetify
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import 'vuetify/styles'

// Components
import App from './App.vue'
import router from './router'


const vuetify = createVuetify({
    components,
    directives,
})

// createApp(App).mount('#app')
// createApp(App).use(vuetify).mount('#app')
createApp(App).use(router).use(vuetify).mount('#app')

import { createApp } from 'vue';

// Vuetify
import '@mdi/font/css/materialdesignicons.css'; // Ensure you are using css-loader
import 'unfonts.css';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import 'vuetify/styles';

// Components
import App from './App.vue';
import router from './router';


const vuetify = createVuetify({
    components,
    directives,
    ssr: true,
    icons: {
        defaultSet: 'mdi', // This is already the default value - only for display purposes
    },
})

// createApp(App).mount('#app')
// createApp(App).use(vuetify).mount('#app')
createApp(App).use(router).use(vuetify).mount('#app')

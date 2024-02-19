import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "./router";

// plugins
import PrimeVue from "primevue/config";
import "primevue/resources/themes/aura-dark-purple/theme.css";
import ToastService from "primevue/toastservice";

const app = createApp(App);

app.use(PrimeVue);
app.use(ToastService);
app.use(router);

app.mount("#app");

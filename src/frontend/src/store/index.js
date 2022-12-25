import { createStore } from "vuex";
import createPersistedState from "vuex-persistedstate";
import animation from "./modules/animation";

export default createStore({
    plugins: [createPersistedState()],
    modules: {
        animation: animation,
    }
})

import { createWebHistory, createRouter } from "vue-router";
import { RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        alias: "/tutorials",
        name: "tutorials",
        component: () => import("./components/TutorialsList.vue"),
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
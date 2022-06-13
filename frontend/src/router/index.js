import { createRouter, createWebHistory } from "vue-router";

import AllEvents from "@/components/AllEvents/AllEvents.vue";
import ByDelegator from "@/components/ByDelegator/ByDelegator.vue";

const routes = [
  {
    path: "/",
    component: AllEvents,
  },
  {
    path: "/by_delegator",
    component: ByDelegator,
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
});

export default router;

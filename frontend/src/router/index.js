import { createRouter, createWebHistory } from "vue-router";

import AllEvents from "@/components/AllEvents/AllEvents.vue";
import ByValidators from "@/components/ByValidators/ByValidators.vue";

const routes = [
  {
    path: "/",
    component: AllEvents,
  },
  {
    path: "/by_validators",
    component: ByValidators,
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
});

export default router;

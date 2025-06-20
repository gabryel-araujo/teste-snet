export default defineNuxtRouteMiddleware((to) => {
  if (import.meta.client) {
    const user = localStorage.getItem("auth-token");
    if (!user) {
      return navigateTo("/");
    }
  }
});

import { createRouter, createWebHistory } from 'vue-router'
import goToLogin from '../views/Login.vue'
import goToDashboard from '../views/Dashboard.vue'
import goToCalendar from '../views/Calendar.vue'
import goToUserInfo from '../views/UserInfo.vue'
import goToTeamsInfo from '../views/TeamsInfo.vue'
import goToRegister from '../views/Registration.vue'

const routes = [
  { path: '/', component: goToLogin }, // Login screen 
  { path: '/register', component: goToRegister }, // Register screen
  { path: '/dashboard', component: goToDashboard }, // Dashboard screen
  { path: '/teaminfo', component: goToTeamsInfo }, // Teams screen
  { path: '/userinfo', component: goToUserInfo }, // User screen
  { path: '/calendar', component: goToCalendar }, // Calendar screen
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router

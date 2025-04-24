<template>
  <div class="flex h-screen w-screen bg-black text-white overflow-hidden">
    <!-- Left navigation bar -->
    <NavigationBar 
      activeSection="userInfo"
      @back="goBack"
      @navigate="navigateTo"
      @toggle-settings="toggleSettings"
      @toggle-info="toggleInfo"
    />
   
    <!-- Main content area -->
    <div class="flex-1 flex flex-col">
      <!-- Header bar -->
      <div class="bg-black pt-5 pb-3 px-8 border-b border-white/50 ">
        <div class="flex items-center">
          <h1 class="text-5xl font-semibold text-white font-adlam">ExeCute</h1>
        </div>
      </div>
      <!-- Points -->
      <div class="text-white text-xs mr-5 mt-2 px-4 py-1 ml-auto bg-infoBg rounded-xl border-borderColor border-2 border-solid ">
        <h1 class="">Points: 1280</h1>
      </div>
     
      <div class="flex flex-col bg-black items-center justify-center pt-[2rem]">
        <div class="w-64 h-64 rounded-full bg-purple-50 flex items-center justify-center overflow-hidden mb-4">
          <!-- Content of the photo -->
        </div>

        <div class="flex px-8 py-4 bg-infoBg items-center justify-center text-center rounded-3xl border-borderColor border-2 border-solid">
          <h1 class="text-[48px]" v-if="userData">{{ userData.username }}</h1>
          <h1 class="text-[48px]" v-else>Loading...</h1>
        </div>
        <div class="px-4 mt-[-1px] bg-infoBg text-center rounded-3xl border-borderColor border-2 border-solid">
          <h1 class="text-[18px]">Art director</h1>
        </div>

        <div class="space-y-2 mt-4 w-full max-w-xl px-6 py-12 bg-infoBg bg-white-200 rounded-xl border-borderColor border-2 border-solid relative">
          <!-- Dots/settings -->
          <div class="absolute top-3 right-6 cursor-pointer">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" class="text-gray-400">
              <circle cx="3" cy="12" r="1.5"></circle>
              <circle cx="11" cy="12" r="1.5"></circle>
              <circle cx="19" cy="12" r="1.5"></circle>
            </svg>
          </div>
          <!-- Email -->
          <div class="w-full bg-fillingInfo rounded-xl px-6 py-2 flex relative">
            <span class="text-left absolute text-white">Email: </span>
            <span class="text-center w-full text-white">iwona@gmail.com</span>
          </div>
          <!-- Phone Number -->
          <div class="w-full bg-fillingInfo rounded-xl px-6 py-2 flex relative">
            <span class="text-left absolute text-white">Phone: </span>
            <span class="text-center w-full text-white">000-000-000</span>
          </div>
          <!-- Birth date -->
          <div class="w-full bg-fillingInfo rounded-xl px-6 py-2 flex relative">
            <span class="text-left absolute text-white">Birth date: </span>
            <span class="text-center w-full text-white">05-09-1994</span>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Dialog settings & design -->
  <SettingsDialog v-model:show="openSettings" @navigate="navigateTo"/>
  <!-- Info Dialog -->
  <InfoDialog v-model:show="openInfo"/>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { ref, onMounted } from 'vue'
import NavigationBar from './NavigationBar.vue'
import SettingsDialog from './PresetsDialogs/SettingsDialog.vue'
import InfoDialog from './PresetsDialogs/InfoDialog.vue'

const router = useRouter()
const userData = ref(null)
const openSettings = ref(false)
const openInfo = ref(false)


//navigation using NavigationBar Preset
const navigateTo = (section) => {
  if (section === 'dashboard') {
    router.push('/dashboard')
  } else if (section === 'teamsInfo') {
    router.push('/teamsInfo')
  } else if (section === 'userInfo') {
    router.push('/userInfo')
  } else if (section === 'calendar') {
    router.push('/calendar')
  } else if (section === 'logout') {
    router.push('/')
  }
}

const toggleSettings = () => {
  openSettings.value = !openSettings.value
}

const toggleInfo = () => {
  openInfo.value = !openInfo.value
}

onMounted(async () => {
  try {
    const validateResponse = await fetch('http://localhost:8437/api/v1/validate', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      },
      credentials: 'include'
    })
    
    if (validateResponse.ok) {
      const validationData = await validateResponse.json()
      const currentUsername = validationData.user
      
      const usersResponse = await fetch('http://localhost:8437/api/v1/user', { 
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        },
        credentials: 'include'
      })
      
      if (usersResponse.ok) { 
        const users = await usersResponse.json()
        userData.value = users.find(user => user.username === currentUsername)
      } else {
        console.error('Failed to fetch user data')
      }
    } else {
      console.error('Session validation failed')
    }
  } catch(error) {
    console.error('Fetching user data failed:', error)
  }
})
</script>
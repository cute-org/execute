<template>
  <div class="flex h-screen w-screen bg-black text-white overflow-hidden">
   <!-- Left navigation bar -->
   <NavigationBar 
      activeSection="teamsInfo"
      @back="goBack"
      @navigate="navigateTo"
      @toggle-settings="toggleSettings"
      @toggle-info="toggleInfo"
    />
   
   <!-- Main content area -->
   <div class="flex-1 flex flex-col min-h-screen overflow-y-auto">
     <!-- Header bar -->
     <div class="bg-black pt-5 pb-3 px-8 mb-20 border-b border-white/50 ">
       <div class="flex items-center">
         <h1 class="text-5xl font-semibold text-white font-adlam">ExeCute</h1>
       </div>
     </div>

     <!-- Main content area -->
      <div class="flex justify-center w-full">
        <div 
          @click="openTeamDialog" 
          class="space-y-2 w-full max-w-[30rem] px-6 py-4 bg-borderColor rounded-xl border-borderColor border-2 border-solid flex justify-center cursor-pointer hover:opacity-90 transition-opacity"
        >
            <span class="text-6xl">{{ teamData.name }}</span> 
        </div>
      </div>
      <!-- Members, meeting, scoreboard -->
      <!-- Members section -->
     <div class="flex flex-col items-center lg:flex-row lg:justify-center w-full px-4 bg-black items-start">
      <div class="space-y-2 m-16 w-full max-w-md px-6 py-8 bg-infoBg rounded-xl border-borderColor border-2 border-solid relative">
        <!-- Headline -->
        <div class="flex justify-center pb-8">
          <h1 class="text-white text-3xl font-bold">Members</h1>
        </div>

        <!-- Members -->
        <div class="w-full bg-fillingInfo rounded-xl px-4 py-3 flex items-center space-x-4">
          <div class="w-6 h-5 rounded-full bg-purple-50"> </div>
          <div class="flex justify-between w-full text-white text-sm">
            <span class="">Member 1</span>
            <span class="text-right">Project Leader</span>
          </div>
        </div>
      </div>
      <!-- Next meeting section -->
      <div class="space-y-2 m-16 w-full max-w-md px-6 py-8 bg-infoBg rounded-xl border-borderColor border-2 border-solid relative">
        <div class="flex justify-center pb-4">
          <h1 class="text-white text-4xl font-bold">Next meeting:</h1>
        </div>
        <!-- Placeholder date -->
        <div class="flex justify-center">
          <h1 class="text-white text-3xl">{{ teamData.meeting }}</h1>
        </div>
     </div>
     <!-- Scoreboard section -->
     <div class="space-y-2 m-16 w-full max-w-md px-6 py-8 bg-infoBg bg-white-200 rounded-xl border-borderColor border-2 border-solid relative">
      <!-- Headline -->
       <!-- We should delete scoreboard tbh -->
      <div class="flex justify-center pb-8">
          <h1 class="text-white text-3xl font-bold">Scoreboard</h1>
        </div>
        <div class="w-full bg-fillingInfo rounded-xl px-4 py-3 flex items-center space-x-4">
          <div class="flex justify-between w-full text-white text-sm">
            <span class="">Team 1 </span>
            <span class="text-right">500 points</span>
          </div>
        </div>
      </div>
    </div>
   </div>
 </div>

 <!-- Dialog settings & design -->
 <SettingsDialog v-model:show="openSettings" @navigate="navigateTo"/>
 <!-- Info Dialog -->
 <InfoDialog v-model:show="openInfo" />
 <!-- Team Dialog -->
 <TeamDialog v-model:show="showTeamDialog" />
</template>

<!-- Setting up router navigation  -->
<script lang="ts" setup>
  import { useRouter } from 'vue-router'
  import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue'
  import { ref, watch, onMounted } from 'vue'
  import NavigationBar from './NavigationBar.vue'
  import SettingsDialog from './PresetsDialogs/SettingsDialog.vue'
  import InfoDialog from './PresetsDialogs/InfoDialog.vue'
  import TeamDialog from './PresetsDialogs/TeamDialog.vue'
  import { fetchTeamInfo, teamData } from './PresetsScripts/GroupInfo'

  const router = useRouter()
  //navigation using NavigationBar Preset
  const navigateTo = (section) => {
    if (section === 'dashboard') {
      router.push('/dashboard')
    } else if (section === 'teamsinfo') {
      router.push('/teamsinfo')
    } else if (section === 'userInfo') {
      router.push('/userInfo')
    } else if (section === 'calendar') {
      router.push('/calendar')
    } else if (section === 'logout') {
      router.push('/')
    }
  }
  
  const openSettings = ref(false)
  const openInfo = ref(false)
  const showTeamDialog = ref(false)

  const toggleSettings = () => {
    openSettings.value = !openSettings.value
  }

  const toggleInfo = () => {
    openInfo.value = !openInfo.value
  } 
  
  const openTeamDialog = () => {
    showTeamDialog.value = true
  }
  
  onMounted(() => {
    fetchTeamInfo()
  })
</script>
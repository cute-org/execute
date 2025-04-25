<template>
  <div class="flex h-screen w-screen bg-black text-white overflow-hidden">
   <!-- Left navigation bar -->
<<<<<<< Updated upstream
    
   <div class="w-16 h-full bg-navBar flex flex-col border-r border-gray-800">
     <!-- Back button -->
     <div class="h-28 w-16 flex items-center justify-center border-b border-white-800" >
       <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
         <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
       </svg>
      </div>

     <!-- Menu buttons -->
       <button @click = "goToDashboard">
          <div class="h-16 w-16 flex items-center justify-center border-b border-white-800">
             <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
           </svg>
         </div>
       </button>

     <button @click= "goToTeamsInfo">
       <div class="h-16 w-16 flex items-center justify-center border-b border-white-800">
         <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
         </svg>
       </div>
     </button>
     
     <button @click= "goToUserInfo">
       <div class="h-16 w-16 flex items-center justify-center border-b border-white-800">
         <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
         </svg>
       </div>
     </button>

     <!-- Current active -->
     <button>
       <div class="h-16 w-16 flex items-center justify-center border-b border-white-800 bg-activeWindowNavBar">
         <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
         </svg>
       </div>
     </button>
     
     
     <!-- Spacer -->
     <div class="flex-grow"></div>
     
     <!-- Bottom buttons -->
     <div class="h-16 w-16 flex items-center justify-center border-t border-white-800">
        <button @click = "openSettings = !openSettings">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
        </button>
      </div>
      <div class="h-16 w-16 flex items-center justify-center border-t border-white-800">
        <button @click = "openInfo = !openInfo">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </button>
      </div>
   </div>
=======
   <NavigationBar 
      activeSection="calendar"
      @back="goBack"
      @navigate="navigateTo"
      @toggle-settings="toggleSettings"
      @toggle-info="toggleInfo"
    />
>>>>>>> Stashed changes

   <!-- Main content area -->
   <div class="flex-1 flex flex-col">
     <!-- Header bar -->
     <div class="bg-black pt-5 pb-3 px-8 border-b border-white/50 ">
       <div class="flex items-center">
         <h1 class="text-5xl font-semibold text-white font-adlam">ExeCute</h1>
       </div>
     </div>
     
     <!-- Main content area -->
     <div class="flex-grow bg-black">
<<<<<<< Updated upstream
        <div class="flex justify-between items-center">
          <div class="flex items-center m-4">

            <!-- PreviousWeek button nav -->
            <button @click="previousWeek" class="bg-black text-white ">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
              </svg>
            </button>

            <!-- Today button nav -->
            <button @click="today" class="bg-borderColor text-white px-4 py-2 rounded-md">
              Today
            </button>

            <!-- NextWeek button nav -->
            <button @click="nextWeek" class="bg-black text-white ">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
              <div class="text-sm ">{{  formatMonthYear(currentWeekStart) }}</div>
          </div>
          <div class="flex items-center">
        </div>
      </div>

      <div class="flex flex-col w-full h-screen">
        <div class="flex bg-black text-white border-t border-gray-700">
          <!-- Days of the week section -->
          <div 
            v-for="(day, index) in weekDays" 
            :key="index" 
            class="flex-1 border-r border-gray-700 text-center py-2"
          >
            <div class="font-bold text-xl">{{ formatWeekday(day) }}</div>
            <div class="font-bold text-xs">{{ formatDayNumber(day) }}</div>
            <div class="text-xs text-gray-300">{{ isToday(day) ? 'Today' : '' }}</div>
          </div>
        </div>
        
        <div class="flex-grow overflow-y-auto">
            <!-- grid for calendar -->
            <div class="grid grid-cols-7 h-full "> 
              <!-- column for a day -->
              <div 
                v-for="(day, dayIndex) in weekDays" 
                :key="dayIndex" 
                class="border-r border-gray-700 h-full "
              >
                <div class="p-2 h-full">

                  <!-- Content placeholder -->
                   <div class = "bg-fillingInfo rounded-2xl my-2 p-2 flex items-center justify-between">
                    <div class="flex items-center">
                      <div 
                          @click="active = !active"
                          :class="['w-4 h-4 rounded-full mr-3', active ? 'bg-green-200' : 'bg-gray-500']"
                        ></div>
                                                
                          <!-- Elements inside  -->
                          <div class="">
                            <div class="text-xl">asfasf</div>
                          </div>
                        </div>
                        <span class="text-white text-sm">123pkt</span>
                    </div>
                    <!-- End of placeholder-->

                </div>
              </div>
            </div>
          </div>
        </div>
    </div>
  </div>
</div>
 
 <TransitionRoot as="template" :show="openSettings">
      <Dialog class="relative z-10" @close="openSettings = false">
        <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
          <div class="fixed inset-0 bg-gray-500/75 transition-opacity" />
        </TransitionChild>
=======
       <!-- Empty content area -->
     </div>
     
   </div>
 </div>
>>>>>>> Stashed changes

    <!-- Dialog settings & design -->
    <SettingsDialog v-model:show="openSettings"  @navigate="navigateTo" />
    <!-- Info Dialog -->
    <InfoDialog v-model:show="openInfo" />
</template>

<!-- Setting up router navigation  -->
<script lang="ts" setup>
  import { useRouter } from 'vue-router'
  import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue'
<<<<<<< Updated upstream
  import { computed, ref } from 'vue';
  import { 
  format, 
  startOfWeek, 
  endOfWeek, 
  eachDayOfInterval, 
  addDays,
  subDays,
  addWeeks,
  subWeeks,
  isToday as dateFnsIsToday,
} from 'date-fns';
//Language for dates, zone
import { enAU } from 'date-fns/locale';
=======
  import { ref } from 'vue';
  import NavigationBar from './NavigationBar.vue'
  import SettingsDialog from './PresetsDialogs/SettingsDialog.vue'
  import InfoDialog from './PresetsDialogs/InfoDialog.vue'
>>>>>>> Stashed changes


   const router = useRouter()
   //Navigation
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

<<<<<<< Updated upstream
   const openSettings = ref(false)
   const openInfo = ref(false)

  //Calendar
   //Current week
   const currentWeekStart = ref(startOfWeek(new Date(), { weekStartsOn: 1}))

   const weekDays = computed(() => {
    return eachDayOfInterval({
      start: currentWeekStart.value,
      end: endOfWeek(currentWeekStart.value, { weekStartsOn: 1})
    })
   })

//Navigations
   //Next week nav
   const nextWeek = () => {
    currentWeekStart.value = addWeeks(currentWeekStart.value, 1)
   }
   //Prev week nav
   const previousWeek = () => {
    currentWeekStart.value = subWeeks(currentWeekStart.value, 1)
   }
   //Today nav
   const today = () => {
    currentWeekStart.value = startOfWeek(new Date(), { weekStartsOn: 1})
   }

// Date functions
   //Month current
    const formatMonthYear = (date: Date) => {
      return format(date, 'MMMM yyyy', { locale: enAU });
    };
    //Days
    const formatDayNumber = (date: Date) => {
      return format(date, 'dd');
    };
    //Current week
    const formatWeekday = (date: Date) => {
      return format(date, 'EEEE', { locale: enAU });
    };
    //Current day
    const isToday = (date: Date) => {
      return dateFnsIsToday(date);
    };

    
    //Delete later when you delete placeholder for elements inside
    const active = ref(false)
=======
  const openSettings = ref(false)
  const openInfo = ref(false)

  const toggleSettings = () => {
    openSettings.value = !openSettings.value
  }

  const toggleInfo = () => {
    openInfo.value = !openInfo.value
  }
   
>>>>>>> Stashed changes
</script>
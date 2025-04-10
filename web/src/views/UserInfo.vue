<template>
  <div class="flex h-screen w-screen bg-black text-white overflow-hidden">
   <!-- Left navigation bar -->
   <div class="w-16 h-full bg-navBar flex flex-col border-r border-gray-800">
     <!-- Back button -->
     <div class="h-28 w-16 flex items-center justify-center border-b border-white-800">
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
     <!-- Current active -->
     <button>
        <div class="h-16 w-16 flex items-center justify-center border-b border-white-800 bg-activeWindowNavBar">
         <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
         </svg>
       </div>
     </button>
     
     <button @click= "goToCalendar">
       <div class="h-16 w-16 flex items-center justify-center border-b border-white-800">
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
       <div class ="w-64 h-64 rounded-full bg-purple-50 flex items-center justify-center overflow-hidden mb-4">
            <!-- Content of te photo -->
       </div>

        <div class ="flex px-8 py-4 bg-infoBg items-center justify-center text-center rounded-3xl border-borderColor border-2  border-solid">
            <h1 class="text-[48px]">HARDCODED Iwona</h1>
        </div>
        <div class ="px-4 mt-[-2px] bg-infoBg text-center rounded-3xl border-borderColor border-2  border-solid">
            <h1 class="text-[18px]">Art director</h1>
        </div>

        <div class ="space-y-2 mt-4 w-full max-w-xl px-6 py-12 bg-infoBg bg-white-200 rounded-xl border-borderColor border-2 border-solid relative">
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
              <span class ="text-left absolute text-white">Email: </span>
              <span class ="text-center w-full text-white">iwona@gmail.com</span>
           </div>
          <!-- Phone Number -->
          <div class="w-full bg-fillingInfo rounded-xl px-6 py-2 flex relative">
            <span class ="text-left absolute text-white">Phone: </span>
            <span class ="text-center w-full text-white">000-000-000</span>
          </div>
          <!-- Birth date -->
          <div class="w-full bg-fillingInfo rounded-xl px-6 py-2 flex relative">
            <span class ="text-left absolute text-white">Birth date: </span>
            <span class ="text-center w-full text-white">05-09-1994</span>
          </div>
        </div>
        
     </div>
   </div>
 </div>

 <!-- Dialog settings & design -->
    <!-- Settings Dialog-->
    <TransitionRoot as="template" :show="openSettings">
      <Dialog class="relative z-10" @close="openSettings = false">
        <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
          <div class="fixed inset-0 bg-gray-500/75 transition-opacity" />
        </TransitionChild>

        <div class="fixed inset-0 z-10 w-screen overflow-y-auto">
          <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
            <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95" enter-to="opacity-100 translate-y-0 sm:scale-100" leave="ease-in duration-200" leave-from="opacity-100 translate-y-0 sm:scale-100" leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95">
              <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
                  <div class="sm:flex sm:items-start">
                    <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                      <DialogTitle as="h3" class="text-base font-semibold text-gray-900">That is a placeholder for settings popup dialog </DialogTitle>
                      <div class="mt-2">
                        <p class="text-sm text-gray-500">12!21(42"/.)/[[]=->..;`~`""]</p>
                      </div>
                    </div>
                  </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
    <!-- Info Dialog -->
    <TransitionRoot as="template" :show="openInfo">
      <Dialog class="relative z-10" @close="openInfo = false">
        <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
          <div class="fixed inset-0 bg-gray-500/75 transition-opacity" />
        </TransitionChild>

        <div class="fixed inset-0 z-10 w-screen overflow-y-auto">
          <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
            <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95" enter-to="opacity-100 translate-y-0 sm:scale-100" leave="ease-in duration-200" leave-from="opacity-100 translate-y-0 sm:scale-100" leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95">
              <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
                  <div class="sm:flex sm:items-start">
                    <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                      <DialogTitle as="h3" class="text-base font-semibold text-gray-900">That site got created by ExeCute Organization</DialogTitle>
                      <div class="mt-2">
                        <p class="text-sm text-gray-500">Thank you for using it, enjoy</p>
                      </div>
                    </div>
                  </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
</template>

<!-- Setting up router navigation  -->
<script lang="ts" setup>
  import { useRouter } from 'vue-router'
  import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue'
  import { ref } from 'vue'

  const router = useRouter()
  // Setting going to Dashboard from Login screen
  const goToCalendar = () => {
    router.push('/calendar')
  }
  const goToTeamsInfo = () => {
    router.push('/teaminfo')
  }
  const goToDashboard = () => {
    router.push('/dashboard')
  }

  const openSettings = ref(false)
  const openInfo = ref(false)
</script>
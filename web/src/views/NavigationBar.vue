<template>
<!-- Left navigation bar -->
<div class="w-16 h-full bg-navBar flex flex-col border-r border-gray-800">
    <!-- Back button -->
    <div class="h-28 w-16 flex items-center justify-center border-b border-white-800">
      <img src="/icon_1.png" alt="" class="h-10 w-10 object-contain" />
    </div>
    
      <!-- Menu buttons -->
      <button @click = "$emit('navigate', 'dashboard')" :class="{ 'bg-activeWindowNavBar': activeSection === 'dashboard' }">
         <div class="h-16 w-16 flex items-center justify-center border-b border-white-800">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </div>
      </button>
   
    <button @click= "$emit('navigate', 'teamsInfo')" :class="{ 'bg-activeWindowNavBar': activeSection === 'teamsInfo' }">
     <div class="h-16 w-16 flex items-center justify-center border-b border-white-800">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
      </div>
    </button>

    <button @click= "$emit('navigate', 'userInfo')" :class="{ 'bg-activeWindowNavBar': activeSection === 'userInfo' }">
      <div class="h-16 w-16 flex items-center justify-center border-b border-white-800">
         <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
         </svg>
       </div>
    </button>

    <button @click= "$emit('navigate', 'calendar')" :class="{ 'bg-activeWindowNavBar': activeSection === 'calendar' }">
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
       <button @click = "$emit('toggle-settings')">
         <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
         </svg>
       </button>
     </div>
     <div class="h-16 w-16 flex items-center justify-center border-t border-white-800">
       <button @click = "$emit('toggle-info')">
           <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
         </svg>
       </button>
     </div>
   </div>
</template>

<script>
export default {
  name: 'NavigationBar',
  props: {
    activeSection: {
      type: String,
      default: '',
      validator: (value) => {
        return ['dashboard', 'teamsInfo', 'userInfo', 'calendar', 'logout'].includes(value);
      }
    }
  }
}
</script>
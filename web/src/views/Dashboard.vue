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
       <!-- Current active -->
        <button>
            <div class="h-16 w-16 flex items-center justify-center border-b border-white-800 bg-activeWindowNavBar">
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
      
      <!-- Team info section -->
      <div class="px-16 py-6">
        <div class="flex items-center">
          <h2 class="text-3xl text-white font-adlam">Team 1</h2>
          <div class="flex ml-4 -space-x-2">
            <!-- Team elements users placeholders  -->
            <div class="w-8 h-8 rounded-full bg-blue-500 border-2 border-black"></div>
            <div class="w-8 h-8 rounded-full bg-red-500 border-2 border-black"></div>
            <div class="w-8 h-8 rounded-full bg-yellow-500 border-2 border-black"></div>
          </div>
          <!-- Number of team members placeholder  -->
          <div class="ml-1 px-2 py-0.5 bg-gray-500 bg-opacity-50 rounded-full text-sm text-white">+21</div>
        </div>
        <!-- Points placeholder -->
        <div class="ml-1 text-[10px] text-white-300 font-adlam">Points: 100/1049</div>
      </div>
      
      <!-- Main content area -->
      <div class="flex-grow bg-black m-12">
        <!-- Dashboard elements  -->
         <div class="flex flex-wrap gap-12 items-start">
            <!-- To-Do -->
              <div class="w-[28rem] flex-none bg-black rounded-3xl p-6 shadow-lg border-2" style="border-color: #3C2650;">
                <!-- Header -->
                <div class="relative flex justify-center items-center mb-6">
                  <h1 class="text-white text-3xl  font-bold">To-Do</h1>
                  <button class="absolute right-0 text-white text-2xl">
                    &bull;&bull;&bull; <!-- Settings dots -->
                  </button>
                </div>

                <!-- Adding tasks design and logic  -->
                <div v-if="toDoTasks.length" class="">
                  <div 
                  v-for = "(item, index) in toDoTasks"
                  :key = "index"
                  class = "bg-fillingInfo rounded-2xl my-2 p-2 flex items-center justify-between"
                  >
                  <div class="flex items-center">
                      <div :class="{'bg-green-500': item.isDone, 'bg-gray-500': !item.isDone}" class="w-4 h-4 rounded-full mr-3" @click="item.isDone = !item.isDone"></div>
                    <!-- Elements inside  -->
                        <div class="">
                            <div class="text-xl">{{ item.name }}</div>
                            <div v-if="item.dueDate" class="text-xs">Date: {{ item.dueDate }}</div> <!-- Show only when it's provided -->
                            </div>
                            
                        </div>
                        <span class="text-white text-sm">{{ item.points }}pkt</span>
                    </div>
                </div>
                  <!-- No tasks yet  -->
                <div v-else class="">No tasks yet. Add one</div>
            <!-- Add task button -->
            <div>
              <button class="flex items-center text-white" @click="openModal('todo')">
                <div class="w-6 h-6 rounded-full bg-transparent border-2 border-white flex items-center justify-center mr-2">
                  <span class="font-bold">+</span>
                </div>
                <span class="text-white text-lg font-medium">Add Tasks</span>
              </button>
            </div>
          </div>


          <!-- In progress -->
          <div class="w-[28rem] bg-black rounded-3xl p-6 shadow-lg border-2" style="border-color: #3C2650;">
                <!-- Header -->
                <div class="relative flex justify-center items-center mb-6">
                  <h1 class="text-white text-3xl  font-bold">In progress</h1>
                  <button class="absolute right-0 text-white text-2xl">
                    &bull;&bull;&bull; <!-- Settings dots -->
                  </button>
                </div>

                <!-- Adding tasks design and logic  -->
                <div v-if="inProgressTasks.length" class="">
                  <div 
                  v-for = "(item, index) in inProgressTasks"
                  :key = "index"
                  class = "bg-fillingInfo rounded-2xl my-2 p-2 flex items-center justify-between"
                  >
                  <div class="flex items-center">
                      <div :class="{'bg-green-500': item.isDone, 'bg-gray-500': !item.isDone}" class="w-4 h-4 rounded-full mr-3" @click="item.isDone = !item.isDone"></div>
                    <!-- Elements inside  -->
                        <div class="">
                            <div class="text-xl">{{ item.name }}</div>
                            <div v-if="item.dueDate" class="text-xs">Date: {{ item.dueDate }}</div> <!-- Show only when it's provided -->
                            </div>
                            
                        </div>
                        <span class="text-white text-sm">{{ item.points }}pkt</span>
                    </div>
                </div>
                  <!-- No tasks yet  -->
                <div v-else class="">No tasks yet. Add one</div>

            <!-- Add task button -->
            <div>
              <button class="flex items-center text-white" @click="openModal('inProgress')">
                <div class="w-6 h-6 rounded-full bg-transparent border-2 border-white flex items-center justify-center mr-2">
                  <span class="font-bold">+</span>
                </div>
                <span class="text-white text-lg font-medium">Add Tasks</span>
              </button>
            </div>
          </div>


          <!-- Completed -->
          <div class="w-[28rem] bg-black rounded-3xl p-6 shadow-lg border-2" style="border-color: #3C2650;">
                <!-- Header -->
                <div class="relative flex justify-center items-center mb-6">
                  <h1 class="text-white text-3xl font-bold">Completed</h1>
                  <button class="absolute right-0 text-white text-2xl">
                    &bull;&bull;&bull; <!-- Settings dots -->
                  </button>
                </div>

                <!-- Adding tasks design and logic  -->
                <div v-if="completedTasks.length" class="">
                  <div 
                  v-for = "(item, index) in completedTasks"
                  :key = "index"
                  class = "bg-fillingInfo rounded-2xl my-2 p-2 flex items-center justify-between"
                  >
                  <div class="flex items-center">
                      <div :class="{'bg-green-500': item.isDone, 'bg-gray-500': !item.isDone}" class="w-4 h-4 rounded-full mr-3" @click="item.isDone = !item.isDone"></div>
                    <!-- Elements inside  -->
                        <div class="">
                            <div class="text-xl">{{ item.name }}</div>
                            <div v-if="item.dueDate" class="text-xs">Date: {{ item.dueDate }}</div> <!-- Show only when it's provided -->
                            </div>
                            
                        </div>
                        <span class="text-white text-sm">{{ item.points }}pkt</span>
                    </div>
                </div>
                  <!-- No tasks yet  -->
                <div v-else class="">No tasks yet. Add one</div>

            <!-- Add task button -->
            <div>
              <button class="flex items-center text-white" @click="openModal('completed')">
                <div class="w-6 h-6 rounded-full bg-transparent border-2 border-white flex items-center justify-center mr-2">
                  <span class="font-bold">+</span>
                </div>
                <span class="text-white text-lg font-medium">Add Tasks</span>
              </button>
            </div>
          </div>
        </div>
    </div>
  </div>
</div>

  <!-- Task Modal -->
  <div v-if="isModalOpen" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-zinc-900 text-white p-6 rounded-xl w-full max-w-md space-y-4">
        <h2 class="text-xl font-semibold">Add New Task</h2>
        <div>
          <label class="block mb-1">Task Name</label>
          <input v-model="task.name" type="text" class="w-full p-2 rounded bg-zinc-800 border border-zinc-700" />
        </div>
        <div>
          <label class="block mb-1">Points</label>
          <input v-model="task.points" type="number" class="w-full p-2 rounded bg-zinc-800 border border-zinc-700" />
        </div>
        <div>
          <label class="block mb-1">Due Date</label>
          <input v-model="task.dueDate" type="date" class="w-full p-2 rounded bg-zinc-800 border border-zinc-700" /> <!-- input adds calendar on the right btw -->
        </div>
        <div class="flex justify-end space-x-2">
          <button @click="closeModal" class="px-4 py-2 rounded bg-gray-600 hover:bg-gray-500">Cancel</button>
          <button @click="saveTask" class="px-4 py-2 rounded bg-blue-600 hover:bg-blue-500">Save</button>
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
    import { ref } from 'vue';

    const router = useRouter()
    // Setting going to Dashboard from Login screen
    const goToCalendar = () => {
    router.push('/calendar')
    }
    const goToUserInfo = () => {
    router.push('/userinfo')
    }
    const goToTeamsInfo = () => {
    router.push('/teaminfo')
    }

    const openSettings = ref(false)
    const openInfo = ref(false)

    //Adding tasks logic
    const isModalOpen = ref(false) //For showing add tasks dialog
    

    function openModal(listType: string) {
      activeTaskList.value = listType
      isModalOpen.value = true
    }
    function closeModal() {
      isModalOpen.value = false
    }

    const toDoTasks = ref<Array<{name: string; points: number; dueDate: string, isDone: Boolean }>>([])
    const inProgressTasks = ref<Array<{name: string; points: number; dueDate: string, isDone: Boolean }>>([])
    const completedTasks = ref<Array<{name: string; points: number; dueDate: string, isDone: Boolean }>>([])
    const activeTaskList = ref('')

    const task = ref({
      name: '',
      points: 0, //Number
      dueDate: '',
      isDone: false
    })
    
    function saveTask() {
      //Placeholder logic for adding in list
      if (!task.value.name.trim()) {
        alert('Task name is required')
        return
      }
      //Logic for backend here
      // Adding to specific list
      if (activeTaskList.value == 'todo') {
        toDoTasks.value.push({...task.value})
      } else if (activeTaskList.value == 'inProgress') {
        inProgressTasks.value.push({...task.value})
      } else if (activeTaskList.value == 'completed') {
        completedTasks.value.push({...task.value})
      }

      closeModal()
    }
</script>
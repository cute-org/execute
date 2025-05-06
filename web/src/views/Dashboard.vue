<template>
   <div class="flex h-screen w-screen bg-black text-white overflow-hidden">
    <!-- Left navigation bar -->
    <NavigationBar 
      activeSection="dashboard"
      @navigate="navigateTo"
      @toggle-settings="toggleSettings"
      @toggle-info="toggleInfo"
    />

    <!-- Main content area -->
    <div class="flex-1 flex flex-col min-h-screen overflow-y-auto">
      <!-- Header bar -->
      <div class="bg-black pt-5 pb-3 px-8 border-b border-white/50 ">
        <div class="flex items-center">
          <h1 class="text-5xl font-semibold text-white font-adlam">ExeCute</h1>
        </div>
      </div>
      
      <!-- Team info section -->
      <div class="px-16 py-6">
        <div class="flex items-center">
          <h2 class="text-3xl text-white font-adlam">{{ teamData.name }}</h2>
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
                            <button @click="openTaskSettings(item, 'todo', index)
                            ">

                          <!-- Elements inside  -->
                              <div class="text-left">
                                  <div class="text-xl">{{ item.name.trim() }}</div>
                                  <div v-if="item.dueDate" class="text-xs">Date: {{ item.dueDate }}</div> <!-- Show only when it's provided -->
                                  </div>
                              </button>
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
                  <span class="font-sans font-bold">+</span>
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
                            <button @click="openTaskSettings(item, 'todo', index)
                            ">

                          <!-- Elements inside  -->
                              <div class="text-left">
                                  <div class="text-xl">{{ item.name.trim() }}</div>
                                  <div v-if="item.dueDate" class="text-xs">Date: {{ item.dueDate }}</div> <!-- Show only when it's provided -->
                                  </div>
                              </button>
                        </div>
                        <span class="text-white text-sm">{{ item.points || '0'}} pkt</span>
                    </div>
                </div>
                  <!-- No tasks yet  -->
                <div v-else class="">No tasks yet. Add one</div>

            <!-- Add task button -->
            <div>
              <button class="flex items-center text-white" @click="openModal('inProgress')">
                <div class="w-6 h-6 rounded-full bg-transparent border-2 border-white flex items-center justify-center mr-2">
                  <span class="font-sans font-bold">+</span>
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
                            <button @click="openTaskSettings(item, 'todo', index)
                            ">
                          <!-- Elements inside  -->
                              <div class="text-left">
                                  <div class="text-xl">{{ item.name.trim() }}</div>
                                  <div v-if="item.dueDate" class="text-xs">Date: {{ item.dueDate }}</div> <!-- Show only when it's provided -->
                                </div>
                            </button>
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
                  <span class="font-sans font-bold">+</span>
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
      <div class="border-2 border-solid border-borderColor bg-infoBg text-white p-6 rounded-xl w-full max-w-md space-y-4">
        <h2 class="text-2xl text-center font-semibold">Add Task</h2>
        <div>
          <label class="block mb-1">Task Name</label>
          <input v-model="task.name" type="text" class="w-full p-2 rounded-xl bg-fillingInfo border border-zinc-700" />
        </div>
        <div>
          <label class="block mb-1">Description</label>
          <input v-model="task.description" type="text" class="w-full p-2 rounded-xl bg-fillingInfo border border-zinc-700" />
        </div>
        <div></div>
        <div>
          <label class="block mb-1">Points</label>
          <input v-model="task.points" type="number" class="w-full p-2 rounded-xl bg-fillingInfo border border-zinc-700" />
        </div>
        <div>
          <label class="block mb-1">Due Date</label>
          <input v-model="task.dueDate" type="datetime-local" class="w-full p-2 rounded-xl bg-fillingInfo border border-zinc-700" /> <!-- input type adds calendar on the right btw -->
        </div>
        <div class="flex justify-end space-x-2">
          <button @click="closeModal" class="px-4 py-2 rounded bg-gray-600 hover:bg-gray-500">Cancel</button>
          <button @click="saveTask" class="px-4 py-2 rounded bg-blue-600 hover:bg-blue-500">Save</button>
        </div>
      </div>
    </div>

    <!-- Settings Modal -->
    <div v-if="isTaskSettingOpen" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
  <div class="border-2 border-solid border-borderColor bg-infoBg text-white p-6 rounded-xl w-full max-w-md space-y-4">
    <h2 class="text-2xl text-center font-semibold">{{ selectedTask?.name }}</h2>
    
    <div>
      <h3 class="font-medium text-lg mb-1">Description:</h3>
      <!-- If no description 'No description provided'-->
      <p class="text-gray-300">{{ selectedTask?.description || 'No description provided' }}</p>
    </div>
    
    <div>
      <h3 class="font-medium text-lg mb-1">Points:</h3>
      <!-- If no points 'No points set'-->
      <p class="text-gray-300">{{ selectedTask?.points || 'No points set' }} </p>
    </div>
    
    <div>
      <h3 class="font-medium text-lg mb-1">Due Date:</h3>
      <!-- If no date 'No due date set'-->
      <p class="text-gray-300">{{ selectedTask?.dueDate || 'No due date set' }}</p>
    </div>
    
    <div class="flex justify-between space-x-2 mt-4">
      <button @click="deleteTask" class="px-4 py-2 rounded bg-red-600 hover:bg-red-500">Delete Task</button>
      <button @click="closeTaskSettings" class="px-4 py-2 rounded bg-gray-600 hover:bg-gray-500">Close</button>
    </div>
  </div>
</div>

    <!-- Dialog settings & design -->
    <SettingsDialog v-model:show="openSettings"  @navigate="navigateTo"/>
    <!-- Info Dialog -->
    <InfoDialog v-model:show="openInfo" />
</template>


<!-- Setting up router navigation  -->
<script lang="ts" setup>
    import { useRouter } from 'vue-router'
    import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue'
    import { onMounted, ref } from 'vue';
    import NavigationBar from './NavigationBar.vue'
    import SettingsDialog from './PresetsDialogs/SettingsDialog.vue'
    import InfoDialog from './PresetsDialogs/InfoDialog.vue'
    import { fetchTeamInfo, teamData } from './PresetsScripts/GroupInfo'


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

    const openSettings = ref(false)
    const openInfo = ref(false)

    const toggleSettings = () => {
      openSettings.value = !openSettings.value
    }

    const toggleInfo = () => {
      openInfo.value = !openInfo.value
    } 

    //Adding tasks logic
    const isModalOpen = ref(false) //For showing add tasks dialog
    

    function openModal(listType: string) {
      activeTaskList.value = listType
      isModalOpen.value = true
    }
    
    function closeModal() {
      isModalOpen.value = false
    }
    
    interface TaskItem {
      name: string,
      description: string,
      points: number,
      dueDate: string,
      isDone: boolean
    }
    
    const task = ref<TaskItem>({
      name: '',
      description: '',
      points: 0, //Number
      dueDate: '',
      isDone: false
    })

    const toDoTasks = ref<TaskItem[]>([])
    const inProgressTasks = ref<TaskItem[]>([])
    const completedTasks = ref<TaskItem[]>([])
    const activeTaskList = ref('')

    async function fetchTasks() {
      try {
        const response = await fetch('http://localhost:8437/api/v1/task', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
          },
          credentials: 'include',
        })
        if (response.ok) {
          const tasks = await response.json()

          tasks.forEach((task: any) => {
            const taskItem: TaskItem = {
              name: task.name,
              description: task.description,
              points: task.pointsValue,
              dueDate: task.dueDate,
              isDone: false,
            }
              if (task.step === 1) {
              toDoTasks.value.push(taskItem)
              } else if (task.step === 2) {
              inProgressTasks.value.push(taskItem)
              } else if (task.step === 3) {
              completedTasks.value.push(taskItem)
            }
          })
        }
      } catch (error) {
        console.error('Fetching tasks failed:', error)
      }
    }
 

    async function saveTask() {
      if (!task.value.name.trim()) {
        alert('Task name is required')
        return
      }
      
      const stepId = {
        todo: 1,
        inProgress: 2,
        completed: 3
      }

      const load = {
        name: task.value.name.trim(),
        description: task.value.description,
        dueDate: new Date(task.value.dueDate).toISOString(),
        pointsValue: Number(task.value.points),
        step: stepId[activeTaskList.value] || 1
      }

      

    try {
        const respone = await fetch('http://localhost:8437/api/v1/task', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(load),
          credentials: 'include',
        })        
      

      // Adding to specific list
      if (activeTaskList.value == 'todo') {
        toDoTasks.value.push({...task.value})
      } else if (activeTaskList.value == 'inProgress') {
        inProgressTasks.value.push({...task.value})
      } else if (activeTaskList.value == 'completed') {
        completedTasks.value.push({...task.value})
      }
      console.log("Load to server:", load)
      closeModal()
    } catch (error) {
        console.error('Request error:', error);
    }
  } 

    
    //Tasks settings
    const isTaskSettingOpen = ref(false)
    const selectedTask = ref<TaskItem | null>(null)
    const selectedTaskList = ref('')
    const selectedTaskIndex = ref(-1)

  
    function openTaskSettings(task: TaskItem, listType: string, index: number) {
      selectedTask.value = task
      selectedTaskList.value = listType
      selectedTaskIndex.value = index
      isTaskSettingOpen.value = true
    }
    
    function closeTaskSettings() {
      isTaskSettingOpen.value = false
    }

    // Delete task
    function deleteTask() {
      if (selectedTaskList.value === 'todo') {
        toDoTasks.value.splice(selectedTaskIndex.value, 1)
      } else if (selectedTaskList.value === 'inProgress') {
        inProgressTasks.value.splice(selectedTaskIndex.value, 1)
      } else if (selectedTaskList.value === 'completed') {
        completedTasks.value.splice(selectedTaskIndex.value, 1)
      }
      closeTaskSettings()
    }

    onMounted(() => {
      fetchTasks()
      fetchTeamInfo()
    })
</script>
import { ref } from 'vue'

export const usersData = ref([])

export const updateTeamData = (newUsersData) => {
    usersData.value = newUsersData
  }
  
export async function fetchTeamUsersInfo() {
  try {
    const response = await fetch('http://localhost:8437/api/v1/group', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    });
    
    if (response.ok) {
      const data = await response.json();
      usersData.value = data;
      console.log('Fetched team users:', usersData.value)
    } else {
      console.error('Error fetching team info:', response.status);
    }
  } catch (error) {
    console.error('Failed to fetch team users info', error);
  }
}
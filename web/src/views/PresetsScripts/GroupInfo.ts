import { ref } from 'vue'

export const teamData = ref({
    name: 'No data',
    code: '',
    meeting: null,
  });

export const updateTeamData = (newTeamData) => {
    teamData.value = newTeamData
  }
  
export async function fetchTeamInfo() {
  try {
    const response = await fetch('http://localhost:8437/api/v1/group/info', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    });
    
    if (response.ok) {
      const data = await response.json();
      teamData.value = data;
      
      // meeting date
      if (teamData.value.meeting) {
        const meetingDate = new Date(teamData.value.meeting);
        teamData.value.meeting = meetingDate.toLocaleString();
      }
    } else {
      console.error('Error fetching team info:', response.status);
    }
  } catch (error) {
    console.error('Failed to fetch team info', error);
  }
}
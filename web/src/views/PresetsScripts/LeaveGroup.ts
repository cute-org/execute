import { fetchTeamInfo } from "./GroupInfo";
const API_BASE_URL = 'http://localhost:8437/api/v1';

export async function handleLeaveGroup(onSuccess= () => {}){
    try {
        const response = await fetch(`${API_BASE_URL}/group/leave`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
              },
              credentials: 'include',
        });
        if (response.ok) {
            const result = await response.json();
            console.log(result.message);
            await fetchTeamInfo();
        }
        onSuccess()
    } catch (error) {
        console.error('Failed to fetch team users info', error);
    }
};
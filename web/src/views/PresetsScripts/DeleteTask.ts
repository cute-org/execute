export async function deleteTask(taskId) {
  try {
    const response = await fetch('http://localhost:8437/api/v1/task', {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',

       body: JSON.stringify({
        taskId: parseInt(taskId)
      })
    });
    
    if (response.ok) {
      const data = await response.json();
      console.log('Task deleted', data);
      return true;
    } else {
      console.error('Error deleting task:', response.status);
      return false;
    }
  } catch (error) {
    console.error('Failed to deleting task', error);
    return false;
  }
}
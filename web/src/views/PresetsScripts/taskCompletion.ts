export async function toggleCompletion(task) {
    try {
        if (!task.id) {
            console.error('Task ID is missing');
            return;
        }
          
      const response = await fetch('http://localhost:8437/api/v1/task/completion', {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',

        body: JSON.stringify({
          taskId: task.id,
          completed: !task.completed
        })
      });
  
      if (!response.ok) {
        throw new Error(`Failed to update task completion: ${response.status}`);
      }
  
      const result = await response.json();
      task.completed = result.completed;
    } catch (error) {
      console.error('Error updating task completion:', error);
    }
  }
  
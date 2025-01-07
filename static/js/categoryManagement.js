// Edit Popup Logic
document.querySelectorAll('.edit-btn').forEach(button => {
    button.addEventListener('click', () => {
      document.getElementById('edit-popup').classList.remove('hidden');
    });
  });
  document.getElementById('close-edit-popup').addEventListener('click', () => {
    document.getElementById('edit-popup').classList.add('hidden');
  });
  document.getElementById('save-edit').addEventListener('click', () => {
    const categoryName = document.getElementById('category-name').value;
    console.log('Save category name:', categoryName); // Data to send to backend
    document.getElementById('edit-popup').classList.add('hidden');
  });
  
  // Manage Offer Popup Logic
  document.querySelectorAll('.manage-offer-btn').forEach(button => {
    button.addEventListener('click', () => {
      document.getElementById('manage-offer-popup').classList.remove('hidden');
    });
  });
  document.getElementById('close-manage-offer-popup').addEventListener('click', () => {
    document.getElementById('manage-offer-popup').classList.add('hidden');
  });
  document.getElementById('save-manage-offer').addEventListener('click', () => {
    const offerData = {
      name: document.getElementById('offer-name').value,
      percentage: document.getElementById('offer-percentage').value,
      description: document.getElementById('offer-description').value,
      startDate: document.getElementById('offer-start-date').value,
      endDate: document.getElementById('offer-end-date').value
    };
    console.log('Save offer data:', offerData); // Data to send to backend
    document.getElementById('manage-offer-popup').classList.add('hidden');
  });
  document.getElementById('delete-offer').addEventListener('click', () => {
    console.log('Delete offer'); // Add backend call for deleting offer
    document.getElementById('manage-offer-popup').classList.add('hidden');
  });
  
  //Add Category Popup
  
   // Function to toggle modal visibility
   function toggleModal() {
    const modal = document.getElementById('categoryModal');
    modal.classList.toggle('hidden');
  }
  
  // Function to handle adding a category (example logic)
  function addCategory() {
    const categoryName = document.getElementById('categoryName').value;
    if (categoryName) {
      console.log(`Category added: ${categoryName}`);
      // You can add further logic to handle the category addition
      toggleModal(); // Close modal after adding
    } else {
      alert('Please enter a category name');
    }
  }
  
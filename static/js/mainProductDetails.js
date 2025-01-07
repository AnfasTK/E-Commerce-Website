// Add Offer Modal
const addOfferBtn = document.getElementById("addOfferBtn");
const offerModal = document.getElementById("offerModal");
const closeAddOfferModal = document.getElementById("closeAddOfferModal");
const addOfferForm = document.getElementById("addOfferForm");

if (addOfferBtn) {
    addOfferBtn.addEventListener("click", () => {
        offerModal.classList.remove("hidden");
    });
}

closeAddOfferModal.addEventListener("click", () => {
    offerModal.classList.add("hidden");
});

addOfferForm.addEventListener("submit", (e) => {
    e.preventDefault();
    alert("Offer saved successfully!");
    offerModal.classList.add("hidden");
    addOfferForm.reset();
});

// Update/Delete Offer Modal
const updateModal = document.getElementById("updateModal");
const closeUpdateModal = document.getElementById("closeUpdateModal");
const deleteOfferBtn = document.getElementById("deleteOfferBtn");
const updateOfferBtn = document.getElementById("updateOfferBtn");
const updateOfferForm = document.getElementById("updateOfferForm");

// Add event listener to show the modal
updateOfferBtn.addEventListener("click", () => {
    updateModal.classList.remove("hidden"); // Show the modal by removing the 'hidden' class
});

closeUpdateModal.addEventListener("click", () => {
    updateModal.classList.add("hidden");
});

deleteOfferBtn.addEventListener("click", () => {
    const confirmation = confirm("Are you sure you want to delete this offer?");
    if (confirmation) {
        alert("Offer deleted successfully!");
        updateModal.classList.add("hidden");
        updateOfferForm.reset();
    }
});

updateOfferForm.addEventListener("submit", (e) => {
    e.preventDefault();
    const formData = {
        offerName: document.getElementById("updateOfferName").value,
        offerDetails: document.getElementById("updateOfferDetails").value,
        startDate: document.getElementById("updateStartDate").value,
        endDate: document.getElementById("updateEndDate").value,
        amount: document.getElementById("updateAmount").value,
        percentage: document.getElementById("updatePercentage").value,
    };
    console.log("Updated Offer Data:", formData);
    alert("Offer updated successfully!");
    updateModal.classList.add("hidden");
    updateOfferForm.reset();
});

const openModal = document.getElementById("openModal");
const closeDescriptionModal = document.getElementById("closeModal");  // Changed name here
const modal = document.getElementById("descriptionModal");
const addPairButton = document.getElementById("addPair");
const keyValuePairs = document.getElementById("keyValuePairs");

// Open modal
openModal.onclick = () => {
    modal.classList.remove("hidden");
}

// Close modal
closeDescriptionModal.onclick = () => {  // Updated here too
    modal.classList.add("hidden");
}

// Add a new pair (heading and description)
addPairButton.onclick = () => {
    const newPair = document.createElement('div');
    newPair.classList.add('flex', 'mb-4', 'space-x-4');
    newPair.innerHTML = `
        <input type="text" name="heading[]" placeholder="Enter heading" class="p-2 border border-gray-300 rounded w-1/3" required>
        <textarea name="description[]" placeholder="Enter description" class="p-2 border border-gray-300 rounded w-2/3" required></textarea>
    `;
    keyValuePairs.appendChild(newPair);
}

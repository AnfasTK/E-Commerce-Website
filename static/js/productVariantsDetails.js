// Modal Functionality
const openPopup = document.getElementById("openPopup");
const popupModal = document.getElementById("popupModal");
const closePopup = document.getElementById("closePopup");

openPopup.addEventListener("click", () => {
    popupModal.classList.remove("hidden");
});

closePopup.addEventListener("click", () => {
    popupModal.classList.add("hidden");
});

// Add More Specification Functionality
const addSpecBtn = document.getElementById("addSpecBtn");
const specificationsList = document.getElementById("specificationsList");

addSpecBtn.addEventListener("click", () => {
    const newSpec = document.createElement("div");
    newSpec.classList.add("specification-item", "flex", "flex-row", "gap-4", "mt-4");
    newSpec.innerHTML = `
           <div class="flex-1">
               <label for="key" class="block text-gray-700 font-medium">Specification Key</label>
               <input type="text" name="key[]" class="w-full p-2 border border-gray-300 rounded-lg" placeholder="e.g., Material" required>
           </div>
           <div class="flex-1">
               <label for="value" class="block text-gray-700 font-medium">Specification Value</label>
               <input type="text" name="value[]" class="w-full p-2 border border-gray-300 rounded-lg" placeholder="e.g., Cotton" required>
           </div>
       `;
    specificationsList.appendChild(newSpec);
});


// Optional: Close the dropdown when clicking outside
document.addEventListener('click', function (event) {
    const isClickInside = event.target.closest('.group');
    if (!isClickInside) {
        document.querySelectorAll('.options').forEach((dropdown) => {
            dropdown.classList.add('hidden');
        });
    }
});

document.addEventListener('DOMContentLoaded', () => {
    const openPopupBtn = document.getElementById('openUpdatePopup');
    const closePopupBtn = document.getElementById('closeUpdatePopup');
    const popupModal = document.getElementById('updatePopupModal');
    const updateSpecificationsForm = document.getElementById('updateSpecificationsForm');

    // Open popup
    openPopupBtn.addEventListener('click', () => {
        popupModal.classList.remove('hidden');
    });

    // Close popup
    closePopupBtn.addEventListener('click', () => {
        popupModal.classList.add('hidden');
    });

    // Close popup if clicking outside the modal
    popupModal.addEventListener('click', (event) => {
        if (event.target === popupModal) {
            popupModal.classList.add('hidden');
        }
    });

    // Form submission handling
    updateSpecificationsForm.addEventListener('submit', async (event) => {
        event.preventDefault();

        // Create FormData object
        const formData = new FormData(updateSpecificationsForm);

        // Convert FormData to an object
        const data = {
            specification_id: formData.getAll('specification_id[]'),
            specification_key: formData.getAll('specification_key[]'),
            specification: formData.getAll('specification[]')
        };

        try {
            const response = await fetch(updateSpecificationsForm.action, {
                method: 'PATCH', // Change back to POST
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data)
            });

            const responseData = await response.json();

            if (response.ok) {
                popupModal.classList.add('hidden');
                // Optionally, you can add a success message or reload the page
                window.location.reload();
            } else {
                // Handle error
                console.error('Submission failed', responseData);
                alert(responseData.error || 'Failed to update specifications');
            }
        } catch (error) {
            console.error('Error:', error);
            alert('An error occurred while updating specifications');
        }
    });
});

function deleteSpecification(specificationId, productId) {
    // Split the IDs if they're passed as a comma-separated string
    const ids = specificationId.split(',');
    const descId = ids[0];
    const prodId = ids[1] || productId;

    fetch(`/admin/products/variant/specification/delete/${descId}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        }
    }).then(response => {
        if (response.ok) {
            // Remove the specification item from the DOM
            const specificationItem = document.querySelector(`.Specifications-item[data-desc-id="${descId}"]`);
            if (specificationItem) {
                specificationItem.remove();
            }
            // Redirect to product details page
            window.location.href = `/admin/products/variant/detail?variant_id=${prodId}`;
        } else {
            // Handle error cases
            console.error('Failed to delete specification');
        }
    }).catch(error => {
        console.error('Error:', error);
    });
}

document.addEventListener('DOMContentLoaded', function() {
    // Add event listeners to all options buttons
    document.querySelectorAll('.options-btn').forEach((btn) => {
        btn.addEventListener('click', function () {
            // Close any previously open dropdowns
            document.querySelectorAll('.options').forEach(dropdown => {
                if (dropdown !== this.nextElementSibling) {
                    dropdown.classList.add('hidden');
                }
            });

            // Toggle the current dropdown
            const dropdown = this.nextElementSibling;
            dropdown.classList.toggle('hidden');
        });
    });

    // Add global click listener to close dropdowns when clicking outside
    document.addEventListener('click', function(event) {
        const optionsButtons = document.querySelectorAll('.options-btn');
        const optionsMenus = document.querySelectorAll('.options');
        
        let clickedOnOptionsButton = false;
        optionsButtons.forEach(btn => {
            if (btn.contains(event.target)) {
                clickedOnOptionsButton = true;
            }
        });

        if (!clickedOnOptionsButton) {
            optionsMenus.forEach(menu => {
                menu.classList.add('hidden');
            });
        }
    });

    // Add event listeners to all "Change" buttons
    document.querySelectorAll('#openUploadPopup').forEach(button => {
        button.addEventListener('click', function() {
            // Find the closest product ID associated with this image
            const productIdInput = document.getElementById('product-id');
            const productId = this.closest('.group').closest('div').querySelector('#product-id')?.value || productIdInput?.value;
            
            if (productId) {
                document.getElementById('product-id').value = productId;
            }
            
            document.getElementById('imageUploadPopup').classList.remove('hidden');
        });
    });

    // Rest of the existing script remains the same
    let cropper = null;
    let currentImageElement = null;
    let currentFile = null;
    const CROP_WIDTH = 400;
    const CROP_HEIGHT = 400;

    // Close Upload Popup
    function closeUploadPopup() {
        document.getElementById('imageUploadPopup').classList.add('hidden');
        document.getElementById('banner-preview').innerHTML = '';
    }

    function confirmUpload() {
        const previewImg = document.getElementById('banner-preview').querySelector('img');
        const productId = document.getElementById('product-id').value;
        
        if (!previewImg) {
            alert('Please upload an image first');
            return;
        }
    
        // Convert image to file
        fetch(previewImg.src)
            .then(res => res.blob())
            .then(blob => {
                const formData = new FormData();
                formData.append('product_image', blob, 'uploaded-image.png');
                formData.append('image_id', productId);
    
                fetch('/admin/products/variant/image/change', {
                    method: 'POST',
                    body: formData
                })
                    .then(response => response.json())
                    .then(data => {
                        if (data.filename) {
                            alert('Image uploaded successfully: ');
                            // Reload the current page
                            window.location.reload();
                        } else {
                            alert('Upload failed');
                        }
                    })
                    .catch(error => {
                        console.error('Upload error:', error);
                        alert('Upload failed');
                    });
            });
    }

    // Drag and Drop Functionality
    function enableDragAndDrop() {
        const dropArea = document.getElementById('banner-drop-area');
        const input = document.getElementById('banner-input');
        const previewContainer = document.getElementById('banner-preview');

        ['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
            dropArea.addEventListener(eventName, preventDefaults, false);
        });

        function preventDefaults(e) {
            e.preventDefault();
            e.stopPropagation();
        }

        dropArea.addEventListener('drop', (e) => {
            const dt = e.dataTransfer;
            const files = dt.files;
            handleFileUpload(files);
        });

        input.addEventListener('change', (e) => {
            handleFileUpload(e.target.files);
        });
    }

    function handleFileUpload(files) {
        const previewContainer = document.getElementById('banner-preview');
        previewContainer.innerHTML = ''; // Clear previous previews

        if (files.length > 1) {
            alert('Please upload only one image.');
            return;
        }

        const file = files[0];
        if (!file.type.startsWith('image/')) {
            alert('Please upload only image files.');
            return;
        }

        const reader = new FileReader();
        reader.onload = function (e) {
            const preview = document.createElement('div');
            preview.className = 'relative border rounded p-2';
            preview.innerHTML = `
                <img src="${e.target.result}" alt="" class="w-full h-40 object-contain">
                <div class="absolute top-2 right-2 flex gap-2">
                    <button type="button" class="bg-blue-500 text-white p-1 rounded" 
                        onclick="startCrop(this.parentElement.parentElement.querySelector('img'), '${file.name}')">
                        Crop
                    </button>
                </div>
            `;
            previewContainer.appendChild(preview);
        };
        reader.readAsDataURL(file);
    }

    function startCrop(imgElement, fileName) {
        const modal = document.getElementById('cropModal');
        const cropImage = document.getElementById('cropImage');

        currentImageElement = imgElement;
        currentFile = fileName;

        cropImage.src = imgElement.src;
        modal.classList.remove('hidden');

        if (cropper) {
            cropper.destroy();
        }

        cropper = new Cropper(cropImage, {
            aspectRatio: CROP_WIDTH / CROP_HEIGHT,
            viewMode: 1,
            dragMode: 'move',
            cropBoxResizable: true,
            cropBoxMovable: true,
            minContainerWidth: 400,
            minContainerHeight: 400,
            imageSmoothingEnabled: true,
            imageSmoothingQuality: 'high',
            background: false,
            modal: false,
            transparent: true
        });
    }

    function cancelCrop() {
        const modal = document.getElementById('cropModal');
        modal.classList.add('hidden');
        if (cropper) {
            cropper.destroy();
            cropper = null;
        }
    }

    function saveCrop() {
        if (!cropper) return;

        const canvas = cropper.getCroppedCanvas({
            width: CROP_WIDTH,
            height: CROP_HEIGHT,
            imageSmoothingEnabled: true,
            imageSmoothingQuality: 'high'
        });

        canvas.toBlob((blob) => {
            const croppedFile = new File([blob], currentFile, { type: 'image/png' });
            const previewContainer = document.getElementById('banner-preview');

            // Update preview with cropped image
            const imgPreview = previewContainer.querySelector('img');
            imgPreview.src = URL.createObjectURL(croppedFile);

            cancelCrop();
        }, 'image/png', 1.0);
    }

    // Expose functions to global scope
    window.closeUploadPopup = closeUploadPopup;
    window.confirmUpload = confirmUpload;
    window.startCrop = startCrop;
    window.cancelCrop = cancelCrop;
    window.saveCrop = saveCrop;

    // Initialize drag and drop
    enableDragAndDrop();
});
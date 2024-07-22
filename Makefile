# Define variables
TARGET_USER = msarya
TARGET_HOST = hpcfund.amd.com
TARGET_DIR = ~/Dev
REMOTE_DIR = ~/Dev/usethis
TARBALL_NAME = project.tar.gz
SOURCE_DIR = .

# Define sync target
sync: create_tarball copy_tarball extract_and_clean_tarball clean

# Create a tarball of the current directory, excluding the tarball itself
create_tarball:
	tar --exclude=$(TARBALL_NAME) -czf $(TARBALL_NAME) $(SOURCE_DIR)

# Copy the tarball to the remote server
copy_tarball:
	scp $(TARBALL_NAME) $(TARGET_USER)@$(TARGET_HOST):$(TARGET_DIR)

# Extract the tarball and delete it on the remote server
extract_and_clean_tarball:
	ssh $(TARGET_USER)@$(TARGET_HOST) 'mkdir -p $(REMOTE_DIR) && tar -xzf $(TARGET_DIR)/$(TARBALL_NAME) -C $(REMOTE_DIR) && rm -f $(TARGET_DIR)/$(TARBALL_NAME)'

# Clean up the tarball after sync
clean:
	rm -f $(TARBALL_NAME)

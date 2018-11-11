from locust import HttpLocust, TaskSet, task

class UserBehavior(TaskSet):

    @task
    def get_something(self):
        self.client.get("/github")
    
    @task
    def get_health(self):
	self.client.get("/health")


class WebsiteUser(HttpLocust):
    task_set = UserBehavior

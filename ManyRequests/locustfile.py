from locust import HttpLocust, TaskSet, task

class UserBehavior(TaskSet):

    @task
    def get_health(self):
        self.client.get("/users")


class WebsiteUser(HttpLocust):
    task_set = UserBehavior



import type { Notification } from "~~/types";

export const useNotifications = () => {
  const notifications = useState<Notification[]>("notifications", () => []);
  const unreadCount = useState<number>("unread_count", () => 0);
  const isConnected = useState<boolean>("notifications_connected", () => false);

  const { apiFetch } = useApi();
  const auth = useAuth();

  const fetchNotifications = async () => {
    if (!auth.isAuthenticated.value) return;

    try {
      const data = await apiFetch<Notification[]>("/notifications");
      notifications.value = data;
      updateUnreadCount();
    } catch (error) {
      console.error("Failed to fetch notifications:", error);
    }
  };

  const markAsRead = async (id: string | number) => {
    try {
      await apiFetch(`/notifications/${id}/read`, {
        method: "PUT",
      });
      await fetchNotifications();
    } catch (error) {
      console.error("Failed to mark notification as read:", error);
    }
  };

  const markAllAsRead = async () => {
    try {
      await apiFetch("/notifications/all/read", {
        method: "PUT",
      });
      await fetchNotifications();
    } catch (error) {
      console.error("Failed to mark all as read:", error);
    }
  };

  const deleteNotification = async (id: number) => {
    try {
      await apiFetch(`/notifications/${id}`, {
        method: "DELETE",
      });
      await fetchNotifications();
    } catch (error) {
      console.error("Failed to delete notification:", error);
    }
  };

  const updateUnreadCount = () => {
    unreadCount.value = notifications.value.filter((n) => !n.read).length;
  };

  const connectWebSocket = () => {
    if (!auth.isAuthenticated.value || isConnected.value) return;

    // In production, use proper WebSocket connection
    // For now, simulate with polling
    const poll = () => {
      fetchNotifications();
      setTimeout(poll, 30000); // Poll every 30 seconds
    };

    poll();
    isConnected.value = true;
  };

  const disconnectWebSocket = () => {
    isConnected.value = false;
  };

  // Show toast notification
  const showToast = (notification: Notification) => {
    if ("Notification" in window && Notification.permission === "granted") {
      new Notification(notification.title, {
        body: notification.message,
        icon: "/favicon-192x192.png",
      });
    }

    // Also show in-app toast
    // You would integrate with a toast component here
    console.log("New notification:", notification);
  };

  // Request notification permission
  const requestPermission = async () => {
    if ("Notification" in window) {
      const permission = await Notification.requestPermission();
      return permission === "granted";
    }
    return false;
  };

  watch(() => auth.isAuthenticated.value, (authenticated) => {
    if (authenticated) {
      fetchNotifications();
      connectWebSocket();
      requestPermission();
    } else {
      notifications.value = [];
      unreadCount.value = 0;
      disconnectWebSocket();
    }
  }, { immediate: true });

  return {
    notifications: readonly(notifications),
    unreadCount: readonly(unreadCount),
    fetchNotifications,
    markAsRead,
    markAllAsRead,
    deleteNotification,
    requestPermission,
    connectWebSocket,
    disconnectWebSocket,
  };
};

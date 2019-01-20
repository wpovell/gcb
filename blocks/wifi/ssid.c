#include <string.h>
#include <sys/socket.h>
#include <sys/ioctl.h>
#include <linux/wireless.h>
#include <unistd.h>
#include "ssid.h"

char* ssid(char* intr) {
  int sock = socket(AF_INET, SOCK_DGRAM, 0);
  if (sock < 0) {
    return NULL;
  }

  char *essid = malloc(sizeof(char) * (IW_ESSID_MAX_SIZE+1));
  struct iwreq wreq;
  memset(&wreq, 0, sizeof(struct iwreq));
  memcpy(wreq.ifr_name, intr, strlen(intr));
  wreq.u.essid.pointer = essid;
  wreq.u.essid.length = IW_ESSID_MAX_SIZE;
  if (ioctl(sock ,SIOCGIWESSID, &wreq) == -1) {
    close(sock);
    return NULL;
  }
  close(sock);

  essid[wreq.u.essid.length] = '\0';
  return essid;
}

import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';

const formatDateTime = (
  date: string | number,
  format: string = 'DD/MM/YYYY'
) => {
  if (!date) {
    return '';
  }

  return dayjs(date).format(format);
};

const unixDateTime = (date: string | number | Date | any) => {
  if (!date) {
    return '';
  }

  return dayjs(date).unix();
};

const readableDateTime = (date: string | number) => {
  if (!date) {
    return '';
  }

  dayjs.extend(relativeTime);

  return dayjs().to(dayjs(date));
};

export { formatDateTime, unixDateTime, readableDateTime };
